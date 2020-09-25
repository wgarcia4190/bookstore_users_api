package app

import (
	"log"

	"github.com/wgarcia4190/bookstore_users_api/internal/logger"

	"github.com/gin-gonic/gin"
)

// StartApplication starts the HTTP server
func StartApplication() {
	router := createRouter()

	mapUrls(router)

	logger.Info("about to start the application...")

	err := router.Run(":8080")
	if err != nil {
		log.Printf("main: Debug service ended %v", err)
	}
}

// createRouter creates the gin.Engine object.
func createRouter() *gin.Engine {
	return gin.Default()
}
