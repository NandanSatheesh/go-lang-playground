package main

import (
	"fmt"
)

func f(Input string) {
	for _, c := range Input {
		fmt.Println(string(c))
	}
}

func main() {

	go f("EFGH")

	f("ABCD")

	go func(Message string) {
		fmt.Println(Message)
	}("This is a message")

	fmt.Scanln()
	fmt.Println("DONE")
}
