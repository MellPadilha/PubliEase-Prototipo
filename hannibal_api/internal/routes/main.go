package routes

import "github.com/gin-gonic/gin"

// AddRoutes adds routes to the provided superRoute.
// It delegates the responsibility of adding specific routes to other functions.
// Accepts a pointer to a Gin RouterGroup as a parameter.
func AddRoutes(superRoute *gin.RouterGroup) {
	formRoutes(superRoute)
}
