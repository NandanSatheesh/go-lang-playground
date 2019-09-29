package main

import (
	"github.com/NandanSatheesh/go-lang-playground/many-one-many-mappings/handlers"
)

func init() {
	handlers.SetupDatabaseHandlers()
}

func main() {

	r := handlers.GetGinHandler()
	r.GET("/", handlers.HealthCheckEndPoint)
	r.Run()
}
