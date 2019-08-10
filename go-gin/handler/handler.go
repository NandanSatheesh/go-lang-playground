package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingGet(c *gin.Context) {

	c.JSON(http.StatusOK, map[string]string{
		"status": "running. OK",
	})
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"appVersion": "1.0",
	})
}
