package main

import (
	"database/sql"
	"fmt"
	"github.com/NandanSatheesh/go-lang-playground/packages/db"
	_ "github.com/NandanSatheesh/go-lang-playground/packages/handles"
)

var database *sql.DB

func init() {
	database = db.CreateCon()
}

func main() {
	fmt.Println("Check Packages")
	fmt.Println(database.Driver())
}
