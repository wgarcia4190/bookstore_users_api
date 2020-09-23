package app

import (
	"github.com/gin-gonic/gin"
	"github.com/wgarcia4190/bookstore_users_api/controllers/health"
	"github.com/wgarcia4190/bookstore_users_api/controllers/users"
)

// mapUrls maps all controller functions with their respective urls and http methods.
func mapUrls(router *gin.Engine) {
	// health_controller endpoints
	router.GET("/health", health.Check)

	// users_controllers endpoints
	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.Create)
}