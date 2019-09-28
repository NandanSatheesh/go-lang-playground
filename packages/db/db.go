package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

/*Create mysql connection*/
func CreateCon() *sql.DB {
	db, err := sql.Open("postgres", "postgres://rcwhbkhp:YwxQ_ElK7Rm9hVMWz3DkjW0qxdKX9o2h@salt.db.elephantsql.com:5432/rcwhbkhp")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("MySQL db is not connected")
		fmt.Println(err.Error())
	}
	return db
}
