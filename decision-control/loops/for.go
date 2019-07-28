package main

import "fmt"

func main() {

	fmt.Println("Time for some loops in GO!")

	fact := 1
	i := 1
	for i <= 5 {
		fact = fact * i
		i = i + 1
	}

	fmt.Println(fact)

	for k := 1; k <= 10; k++ {
		fact *= k
	}

	fmt.Println(fact)

	i = 1
	for ; ; i++ {

		fmt.Println(i)
		if i%2 == 0 {
			continue
		}

		if i == 101 {
			break
		}

	}
}
