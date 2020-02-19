package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page Hit")
}

func testPostHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test Home Page Hit")
}

func allArticles(w http.ResponseWriter, r *http.Request) {

	articles := Articles{
		Article{
			Title:   "Test",
			Desc:    "Description",
			Content: "Some Content",
		},
		Article{
			Title:   "Test",
			Desc:    "Description",
			Content: "Some Content",
		},
		Article{
			Title:   "Test",
			Desc:    "Description",
			Content: "Some Content",
		},
		Article{
			Title:   "Test",
			Desc:    "Description",
			Content: "Some Content",
		},
		Article{
			Title:   "Test",
			Desc:    "Description",
			Content: "Some Content",
		},
		Article{
			Title:   "Test",
			Desc:    "Description",
			Content: "Some Content",
		},
	}
	json.NewEncoder(w).Encode(articles)
}

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/", testPostHomePage).Methods("POST")
	router.HandleFunc("/all", allArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func main() {
	handleRequests()
}
