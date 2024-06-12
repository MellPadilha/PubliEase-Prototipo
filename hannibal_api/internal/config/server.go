package config

import (
	"github.com/ccesarfp/hannibal/internal/routes"
	"github.com/gin-gonic/gin"
)

// Server struct encapsulates the Gin engine along with environment and port configurations.
type Server struct {
	gin  *gin.Engine
	Env  string
	Port string
}

// NewServer creates and returns a new Server instance.
func NewServer() *Server {
	return &Server{}
}

// SetupServer configures the Gin engine, including setting the environment mode,
// middlewares, and routes.
// It accepts an optional middleware function as a parameter.
func (s *Server) SetupServer(optMiddleware gin.HandlerFunc) {

	// Set production environment if specified
	if s.Env == "PRO" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Creating a new Gin engine instance and assigning to the Server struct
	s.gin = gin.New()

	// Setting up middlewares
	s.setupMiddlewares(s.gin, optMiddleware)

	// Setting up routes
	s.setupRoutes(s.gin)

}

// Up starts the Gin server on the specified port.
// It returns an error if the server fails to start.
func (s *Server) Up() error {
	if err := s.gin.Run(s.Port); err != nil {
		return err
	}
	return nil
}

// setupMiddlewares configures the default middlewares (Logger, Recovery) and the optional middleware
// provided as a parameter. It assigns these middlewares to the given Gin engine instance.
func (s *Server) setupMiddlewares(g *gin.Engine, optMiddleware gin.HandlerFunc) {
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(optMiddleware)
}

// setupRoutes configures the API routes for the given Gin engine instance.
// It adds routes defined in the routes package to the "/api/v1" endpoint.
func (s *Server) setupRoutes(g *gin.Engine) {
	router := g.Group("/api/v1")
	routes.AddRoutes(router)
}
