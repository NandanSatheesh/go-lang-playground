package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func init() {
	database, err := sql.Open("mysql", "root:password@/notes?charset=utf8")
	checkErr(err)

	fmt.Println(database)
}

func getAllItem(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getAllItems",
	})

}

func addItem(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "AddItem",
	})
}

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/getAllItems", getAllItem)

	r.POST("/addItem", addItem)

	r.Run()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
