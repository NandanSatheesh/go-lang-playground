package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type QuotesData []string

func main() {

	file, err := ioutil.ReadFile("quotes.json")

	if err != nil {
		fmt.Println(err)
	}

	var data QuotesData

	_ = json.Unmarshal([]byte(file), &data)

	for i, quote := range data {
		fmt.Println(i, " -> ", quote)
	}
}
