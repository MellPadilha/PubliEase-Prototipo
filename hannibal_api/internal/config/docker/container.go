package docker

type Container struct {
	ID       string
	Image    string
	HostIp   string
	Ports    []string
	Volumes  []string
	Name     string
	EnvVars  []string
	Commands []string
	WorkDir  string
}
