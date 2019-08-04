package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/quotes_data?charset=utf8")
	checkErr(err)

	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(60) + 1

	stmt, err := db.Prepare("SELECT data FROM quotes WHERE id = ?")
	checkErr(err)

	var data string
	defer stmt.Close()

	rows, err := stmt.Query(id)
	checkErr(err)

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&data)
		fmt.Println(data)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
