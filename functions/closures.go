package main

import "fmt"

func main() {

	gen := innerFunctionDemo()
	for i := 0; i <= 10; i++ {
		fmt.Println(gen())
	}

}

func innerFunctionDemo() func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, (a + b)
		return a
	}
}
