package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("empty array ", a)

	a[3] = 100
	fmt.Println("the array now has ", a)

	b := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(b)

	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" ", b[j])
		}
		fmt.Println()
	}

	var fib [15]int

	fib[0] = 0
	fib[1] = 1

	for i := 2; i < 15; i++ {
		fib[i] = fib[i-1] + fib[i-2]
		fmt.Println(fib)
	}

}
