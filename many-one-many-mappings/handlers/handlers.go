package handlers

import (
	"github.com/NandanSatheesh/go-lang-playground/many-one-many-mappings/database"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetupDatabaseHandlers() {
	db = database.GetDatabaseHandler()
}

func GetGinHandler() *gin.Engine {
	return gin.Default()
}

func HealthCheckEndPoint(context *gin.Context) {
	context.JSON(200, gin.H{"message": "Running Fine"})
}
