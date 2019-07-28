package main

import (
	"fmt"
	"strconv"
)

const n int = 1000
const message string = "You shouldn't be doing this"

func main() {
	s := strconv.Itoa(n)
	fmt.Println("You suck - " + s)
	fmt.Println(message)
}
