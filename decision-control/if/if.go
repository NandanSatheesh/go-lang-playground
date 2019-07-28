package main

import "fmt"

func main() {

	if num := 10; false {
		fmt.Println("You suck")
	} else if false {
		fmt.Println("You suck anyway")
	} else {
		for i := 0; i < num; i++ {
			fmt.Println("You are on a different level mate")
		}
	}
}
