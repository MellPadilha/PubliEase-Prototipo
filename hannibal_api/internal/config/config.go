package config

import (
	"github.com/ccesarfp/hannibal/internal/config/cache"
	"github.com/ccesarfp/hannibal/internal/config/docker"
	"sync"
)

// Configuration struct represents the configuration for the application,
// containing server and Docker configurations.
type Configuration struct {
	Server *Server
	Docker *docker.Docker
}

var once sync.Once
var instance *Configuration

// New creates a new Configuration instance by initializing server and Docker configurations.
// It returns a pointer to the Configuration struct and an error, if any occurs during initialization.
func New() *Configuration {
	once.Do(func() {
		// Initialize Configuration struct
		instance = &Configuration{}
	})

	return instance
}

func (c *Configuration) Configure() error {
	// Initialize server configuration
	c.Server = NewServer()

	// Setup Viper configuration
	err := setupViper()
	if err != nil {
		return err
	}

	// Retrieve environment values and populate the Configuration struct
	err = getConfigValues(c)
	if err != nil {
		return err
	}

	// Create cache directory
	ch := cache.New()
	err = ch.CreateCacheDir()
	if err != nil {
		return err
	}

	// Initialize Docker configuration
	c.Docker, err = docker.SetupDocker(c.Docker.Container)
	if err != nil {
		return err
	}

	return nil
}
