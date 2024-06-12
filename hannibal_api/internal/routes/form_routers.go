package routes

import (
	"github.com/ccesarfp/hannibal/internal/controllers"
	"github.com/gin-gonic/gin"
)

// formRoutes defines the routes related to forms.
// It accepts a pointer to a Gin RouterGroup as a parameter.
func formRoutes(superRoute *gin.RouterGroup) {
	// Create a new route group "/form" under the provided superRoute.
	formRouter := superRoute.Group("/form")
	{
		formController := controllers.NewFormController()
		formRouter.POST("/", formController.ValidateApplication)
	}
}
