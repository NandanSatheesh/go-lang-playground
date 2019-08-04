package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Bird struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

type KanyeQuotes struct {
	QuoteString string `json:"quote"`
}

func GetBirdData(w http.ResponseWriter, r *http.Request) {

	pigeon := &Bird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}

	data, _ := json.Marshal(pigeon)

	fmt.Println(string(data))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetRandomQuote(w http.ResponseWriter, r *http.Request) {

	quotesData := KanyeQuotes{
		QuoteString: "likes to do weed",
	}

	fmt.Println(quotesData)

	data, err := json.Marshal(quotesData)

	fmt.Println(err)

	fmt.Println(string(data))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage API Hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/bird", GetBirdData)
	http.HandleFunc("/quote", GetRandomQuote)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
