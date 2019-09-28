package handlers

import (
	"fmt"
	"github.com/NandanSatheesh/go-lang-playground/orm/database"
	"github.com/NandanSatheesh/go-lang-playground/orm/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetupDatabaseHandlers() {
	db = database.GetDatabaseHandler()
	db.AutoMigrate(&models.Person{})
}

func GetGinHandler() *gin.Engine {
	return gin.Default()
}

func HealthCheckEndPoint(context *gin.Context) {
	context.JSON(200, gin.H{"message": "Running Fine"})
}

func GetAllPeople(c *gin.Context) {
	var people []models.Person
	fmt.Println(db.Find(&people).Error)
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}
}

func FindPeopleById(context *gin.Context) {
	id := context.Params.ByName("id")
	var person models.Person
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		context.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		context.JSON(200, person)
	}
}

func AddNewPerson(context *gin.Context) {
	var person models.Person
	context.BindJSON(&person)
	if person.FirstName != "" && person.LastName != "" {
		db.Create(&person)
		context.JSON(200, person)
	} else {
		ErrorMessage := models.ErrorMessage{ErrorMessage: "First Name or Last Name Cannot be empty", ErrorCode: 400}
		context.JSON(400, ErrorMessage)
	}
}

func DeletePerson(context *gin.Context) {
	var person models.Person
	id := context.Params.ByName("id")
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	context.JSON(200, gin.H{"id:" + id: "\n\"status\":\"Deleted\""})
}

func UpdatePerson(context *gin.Context) {
	var person models.Person
	id := context.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		context.AbortWithStatus(404)
		fmt.Println(err)
	}
	context.BindJSON(&person)
	db.Save(&person)
	context.JSON(200, person)
}
