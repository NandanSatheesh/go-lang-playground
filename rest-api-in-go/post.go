package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ToDoListData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func addItem(w http.ResponseWriter, req *http.Request) {

	reqBody, err := ioutil.ReadAll(req.Body)
	checkErr(err)

	fmt.Printf("%s \n", reqBody)
}

func handleRequest() {
	http.HandleFunc("/add", addItem)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
