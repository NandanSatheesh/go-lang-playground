package handlers

import (
	"github.com/NandanSatheesh/go-lang-playground/mini-book-search/database"
	"github.com/NandanSatheesh/go-lang-playground/mini-book-search/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetupDatabaseHandlers() {
	db = database.GetDatabaseHandler()

	db.AutoMigrate(&models.Publisher{}, &models.Author{}, &models.Book{})

	populateDataForDb()
}

func populateDataForDb() {
	var authors models.Author
	var publisher models.Publisher
	var book models.Book

	authors = models.Author{
		AuthorName:        "Chris Anderson",
		AuthorNationality: "USA",
	}

	publisher = models.Publisher{
		PublisherName: "NB",
	}

	db.Create(&authors)

	db.Create(&publisher)

	book = models.Book{
		BookTitle:   "TED Talks",
		BookYear:    2012,
		AuthorId:    authors.ID,
		PublisherId: publisher.ID,
	}

	db.Create(&book)

	authors = models.Author{
		AuthorName:        "Ray Dalio",
		AuthorNationality: "USA",
	}

	publisher = models.Publisher{
		PublisherName: "Simon",
	}

	db.Create(&authors)

	db.Create(&publisher)

	book = models.Book{
		BookTitle:   "Principles - Life and Work",
		BookYear:    2015,
		AuthorId:    authors.ID,
		PublisherId: publisher.ID,
	}

	db.Create(&book)
}

func GetGinHandler() *gin.Engine {
	return gin.Default()
}

func HealthCheckEndPoint(context *gin.Context) {
	context.JSON(200, gin.H{"message": "Running Fine"})
}
