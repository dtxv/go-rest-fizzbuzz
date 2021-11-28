package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health check
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
