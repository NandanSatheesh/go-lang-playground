package main

import "fmt"

func main() {
	fmt.Println(fact(12))
	fmt.Println(fact(5))
	fmt.Println(add(10, 30, 30))
	a, b := 10, 30
	a, b = swap(a, b)
	fmt.Println(a, b)

	array := []int{1, 2, 3, 4, 5}
	fmt.Println(printArrayElements(array...))

}

func printArrayElements(elements ...int) int {
	sum := 0
	for _, i := range elements {
		sum += i
		fmt.Println(i)

	}
	fmt.Println(sum)
	return sum
}

func swap(a, b int) (int, int) {
	return b, a
}

func add(a, b, c int) int {
	return a + b + c
}

func fact(f int) int {
	if f == 0 {
		return 1
	}
	return f * fact(f-1)
}
