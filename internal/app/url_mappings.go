package app

import (
	"github.com/gin-gonic/gin"
	"github.com/wgarcia4190/bookstore_users_api/internal/controllers/health"
	"github.com/wgarcia4190/bookstore_users_api/internal/controllers/users"
)

// mapUrls maps all controller functions with their respective urls and http methods.
func mapUrls(router *gin.Engine) {
	// health_controller endpoints
	router.GET("/health", health.Check)

	// users_controllers endpoints
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
}
