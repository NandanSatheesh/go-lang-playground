package main

import (
	"github.com/NandanSatheesh/go-lang-playground/orm/handlers"
)

func init() {
	handlers.SetupDatabaseHandlers()
}

func main() {

	r := handlers.GetGinHandler()
	r.GET("/people", handlers.GetAllPeople)
	r.GET("/", handlers.HealthCheckEndPoint)
	r.GET("/people/:id", handlers.FindPeopleById)
	r.POST("/people/add", handlers.AddNewPerson)
	r.PUT("/people/:id", handlers.UpdatePerson)
	r.DELETE("/people/:id", handlers.DeletePerson)
	r.Run()
}
