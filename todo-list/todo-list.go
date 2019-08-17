package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

type TodoItem struct {
	Name        string `json: "name"`
	Description string `json: "description"`
}

type TransformedTodo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoListItem struct {
	id          int
	title       string
	description string
}

func init() {
	db, err := sql.Open("mysql", "root:password@/notes?charset=utf8")
	checkErr(err)

	database = db

	fmt.Println(database)
}

func getAllItem(c *gin.Context) {

	rows, err := database.Query("SELECT id, title, description FROM todo_list")
	checkErr(err)
	defer rows.Close()

	var finalList []TransformedTodo

	for rows.Next() {
		item := TodoListItem{}
		rows.Scan(&item.id, &item.title, &item.description)
		transformedTodo := TransformedTodo{
			ID:          item.id,
			Title:       item.title,
			Description: item.description}

		finalList = append(finalList, transformedTodo)
	}

	c.JSON(200, gin.H{
		"data":         finalList,
		"message":      "Success",
		"responseCode": 200,
	})

}

func addItem(c *gin.Context) {

	var item TodoItem
	c.BindJSON(&item)

	fmt.Println(item)

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
