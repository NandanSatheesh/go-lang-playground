package main

import "fmt"

func calculate(x int) (result int) {
	result = x + 2
	return result

}

func main() {

	fmt.Println(calculate(12))
}
