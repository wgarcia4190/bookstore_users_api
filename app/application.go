package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

// StartApplication starts the HTTP server
func StartApplication() {
	router := createRouter()

	mapUrls(router)

	err := router.Run(":8080")
	if err != nil {
		log.Printf("main: Debug service ended %v", err)
	}
}

// createRouter creates the gin.Engine object.
func createRouter() *gin.Engine {
	return gin.Default()
}
