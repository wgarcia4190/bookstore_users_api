package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Check responds with a 200 OK if the service is healthy and ready for traffic.
// It is going to be executed when the user hits this url:
// GET /health
func Check(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
