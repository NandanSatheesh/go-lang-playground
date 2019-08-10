package main

import (
	"github.com/gin-gonic/gin"

	"go-gin/handler"
)

func main() {
	r := gin.Default()

	r.POST("/ping", handler.PingGet)

	r.POST("/appVersion", handler.HealthCheck)

	r.Run()
}
