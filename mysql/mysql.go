package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/notes?charset=utf8")
	checkErr(err)

	fmt.Println(db)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
