package middlewares

import (
	"github.com/ccesarfp/hannibal/internal/config"
	"github.com/gin-gonic/gin"
	"log"
)

// Docker middleware checks the availability of the Docker service.
func Docker(config *config.Configuration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ping the Docker service to check availability
		thereIsAProblem := config.Docker.Ping()
		if thereIsAProblem != nil {
			panic(thereIsAProblem)
		}

		// Preparing docker image and container
		err := config.Docker.PrepareContainer()
		if err != nil {
			log.Panicln(err)
		}

		// Starting container
		err = config.Docker.RunContainer()
		if err != nil {
			log.Panicln(err)
		}

		// Proceed to the next middleware or handler
		c.Next()
	}
}
