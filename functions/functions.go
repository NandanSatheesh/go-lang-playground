package main

import "fmt"

func main() {
	fmt.Println(fact(12))
	fmt.Println(fact(5))
	fmt.Println(add(10, 30, 30))
}

func add(a, b, c int) int {
	return a + b + c
}

func fact(f int) int {
	if f == 0 {
		return 1
	} else {
		return f * fact(f-1)
	}
}
