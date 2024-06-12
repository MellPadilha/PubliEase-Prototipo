package main

import (
	"github.com/ccesarfp/hannibal/internal/config"
	"github.com/ccesarfp/hannibal/internal/config/docker"
	"github.com/ccesarfp/hannibal/internal/middlewares"
	"log"
	"time"
)

var Config *config.Configuration
var startTime = time.Now()

func main() {
	var err error

	// Starting config struct, getting env values and starting docker client
	log.Println("[Application] Configuring server")
	Config = config.New()
	err = Config.Configure()
	if err != nil {
		log.Panicln(err)
	}
	defer func(Docker *docker.Docker) {
		err = Docker.Close()
		if err != nil {

		}
	}(Config.Docker)

	// Verifying if docker is running
	log.Println("[Application] Verifying Docker")
	if err = Config.Docker.Ping(); err != nil {
		log.Panicln(err)
	}

	// Preparing docker image and container
	log.Println("[Application] Pulling and creating container")
	if err = Config.Docker.PrepareContainer(); err != nil {
		log.Panicln(err)
	}

	// Starting container
	log.Println("[Application] Starting container")
	err = Config.Docker.RunContainer()
	if err != nil {
		log.Panicln(err)
	}

	// Executing commands
	log.Println("[Application] Configuring container")
	go func() {
		err = Config.Docker.ExecuteCommands()
		if err != nil {
			log.Panicln(err)
		}
	}()

	// Setting server
	log.Println("[Application] Setting Server")
	Config.Server.SetupServer(middlewares.Docker(Config))

	// Serving server
	log.Println("[Application] Server initialization took", time.Since(startTime))
	if err = Config.Server.Up(); err != nil {
		panic(err)
	}
}
