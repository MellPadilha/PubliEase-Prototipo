package docker

import (
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"io"
	"log"
	"strings"
	"sync"
)

// Docker struct encapsulates a Docker Client.
type Docker struct {
	Client    *client.Client
	Container Container
}

var instance *Docker
var once sync.Once

func New() (*Docker, error) {
	var docker *client.Client
	var err error

	once.Do(func() {
		docker, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

		instance = &Docker{
			Client: docker,
		}
	})

	if err != nil {
		return nil, err
	}

	return instance, nil
}

func newWithContainer(container Container) (*Docker, error) {
	var docker *client.Client
	var err error

	once.Do(func() {
		docker, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

		instance = &Docker{
			Client:    docker,
			Container: container,
		}
	})

	if err != nil {
		return nil, err
	}

	return instance, nil
}

// SetupDocker initializes a Docker Client with environment variables and API version negotiation.
// It returns a pointer to a Docker struct and an error, if any occurs during Client creation.
func SetupDocker(container Container) (*Docker, error) {
	docker, err := newWithContainer(container)
	if err != nil {
		return nil, err
	}

	return docker, nil
}

// Close terminates the Docker Client connection.
// It returns an error if the close operation fails.
func (d *Docker) Close() error {
	err := d.Client.Close()
	if err != nil {
		return err
	}
	return nil
}

// Ping checks the connection to the Docker daemon.
// Returns error if unable to connect to Docker.
func (d *Docker) Ping() error {
	_, err := d.Client.Ping(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// PrepareContainer creates and starts a Docker container based on the provided configuration.
// It first ensures that the required image is available, then creates and starts the container.
func (d *Docker) PrepareContainer() error {
	if d.imageExists() == false {
		err := d.pullImage()
		if err != nil {
			return err
		}
	}

	containers, err := d.getAllContainers()
	if err != nil {
		return err
	}

	if d.containerExists(containers) == false {
		resp, err := d.createContainer()
		if err != nil {
			return err
		}
		d.Container.ID = resp.ID
	}

	if d.Container.ID == "" {
		d.Container.ID = d.getContainerID(containers)
	}

	return nil
}

// pullImage pulls the Docker image specified in the configuration if it is not already present locally.
func (d *Docker) pullImage() error {
	read, err := d.Client.ImagePull(
		context.Background(),
		d.Container.Image,
		image.PullOptions{},
	)
	if err != nil {
		return err
	}
	defer read.Close()

	io.Copy(io.Discard, read)

	return nil
}

// imageExists checks if a Docker image with the specified name exists locally.
func (d *Docker) imageExists() bool {
	images, err := d.Client.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		log.Fatalln("Problem getting images: ", err)
		return false
	}
	for _, img := range images {
		for _, tag := range img.RepoTags {
			if tag == d.Container.Image {
				return true
			}
		}
	}
	return false
}

// getAllContainers retrieves a list of all Docker containers on the system.
func (d *Docker) getAllContainers() ([]types.Container, error) {
	containers, err := d.Client.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	return containers, nil
}

// containerExists checks if a Docker container with the specified name exists.
func (d *Docker) containerExists(containers []types.Container) bool {
	for _, cont := range containers {
		for _, name := range cont.Names {
			if name == "/"+d.Container.Name { // container.Names includes a leading slash
				return true
			}
		}
	}
	return false
}

// createContainer creates a new Docker container based on the provided configuration.
func (d *Docker) createContainer() (*container.CreateResponse, error) {
	resp, err := d.Client.ContainerCreate(
		context.Background(),
		d.prepareConfig(),
		d.prepareHostConfig(),
		nil,
		nil,
		d.Container.Name,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// getContainerID fetches the container ID for a given container name
func (d *Docker) getContainerID(containers []types.Container) string {
	for _, cont := range containers {
		for _, name := range cont.Names {
			if name == "/"+d.Container.Name { // container.Names includes a leading slash
				return cont.ID
			}
		}
	}
	return ""
}

// containerIsRunning checks if a Docker container with the specified name is running.
func (d *Docker) containerIsRunning(containers []types.Container) bool {
	for _, cont := range containers {
		for _, name := range cont.Names {
			if name == "/"+d.Container.Name { // container.Names includes a leading slash
				return true
			}
		}
	}
	return false
}

// prepareConfig prepares the container configuration based on the provided container settings.
func (d *Docker) prepareConfig() *container.Config {
	return &container.Config{
		Image:      d.Container.Image,
		Env:        d.Container.EnvVars,
		WorkingDir: d.Container.WorkDir,
	}
}

// parsePortMapping parses the port mapping in the format "hostPort:containerPort"
func (d *Docker) parsePortMapping(portMapping string) (nat.Port, nat.PortBinding, error) {
	parts := strings.Split(portMapping, ":")
	if len(parts) != 2 {
		return "", nat.PortBinding{}, fmt.Errorf("invalid port mapping format: %s", portMapping)
	}

	containerPort := nat.Port(parts[1] + "/tcp") // Assuming TCP, change if necessary
	hostPort := parts[0]

	return containerPort, nat.PortBinding{
		HostIP:   d.Container.HostIp,
		HostPort: hostPort,
	}, nil
}

// parseVolumeMapping parses the volume mapping in the format "source:destination"
func (d *Docker) parseVolumeMapping(volumeMapping string) (mount.Mount, error) {
	parts := strings.Split(volumeMapping, ":")
	if len(parts) != 2 {
		return mount.Mount{}, fmt.Errorf("invalid volume mapping format: %s", volumeMapping)
	}

	source := parts[0]
	destination := parts[1]

	return mount.Mount{
		Type:   mount.TypeBind,
		Source: source,
		Target: destination,
	}, nil
}

// prepareVolumes prepares the volume configuration for the container
func (d *Docker) prepareVolumes() ([]mount.Mount, error) {
	var mounts []mount.Mount

	for _, volume := range d.Container.Volumes {
		mountVolume, err := d.parseVolumeMapping(volume)
		if err != nil {
			return nil, err
		}
		mounts = append(mounts, mountVolume)
	}

	return mounts, nil
}

// prepareHostConfig prepares the host configuration for the container based on the provided container settings.
func (d *Docker) prepareHostConfig() *container.HostConfig {
	portBindings := make(nat.PortMap)

	for _, port := range d.Container.Ports {
		natPort, portBinding, err := d.parsePortMapping(port)
		if err != nil {
			return nil
		}
		portBindings[natPort] = append(portBindings[natPort], portBinding)
	}

	volumes, err := d.prepareVolumes()
	if err != nil {
		return nil
	}

	return &container.HostConfig{
		PortBindings: portBindings,
		Mounts:       volumes,
	}
}

// RunContainer starts the Docker container specified by the creation response.
func (d *Docker) RunContainer() error {
	err := d.Client.ContainerStart(
		context.Background(),
		d.Container.ID,
		container.StartOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

// ExecuteCommands execute commands inside a docker container.
func (d *Docker) ExecuteCommands() error {
	for _, cmd := range d.Container.Commands {
		_, err := d.Execute(cmd)
		if err != nil {
			return err
		}
	}

	return nil
}

// Execute execute command inside a docker container.
func (d *Docker) Execute(cmd string) (string, error) {
	execConfig := types.ExecConfig{
		Cmd:          []string{"sh", "-c", cmd},
		AttachStdout: true,
		AttachStderr: true,
	}
	log.Println("[Docker] " + cmd)
	execIDResp, err := d.Client.ContainerExecCreate(context.Background(), d.Container.ID, execConfig)
	if err != nil {
		return "", err
	}

	resp, err := d.Client.ContainerExecAttach(context.Background(), execIDResp.ID, types.ExecStartCheck{})
	if err != nil {
		return "", err
	}
	defer resp.Close()

	// Capture the output in a variable
	var buf bytes.Buffer
	io.Copy(&buf, resp.Reader)

	output := string(buf.Bytes())

	return output, nil
}
