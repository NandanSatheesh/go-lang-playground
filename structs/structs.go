package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	person := Person{"Nandan", 22}
	fmt.Println(person)
}
