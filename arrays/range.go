package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}

	var sum int
	sum = 0

	for i, num := range nums {
		fmt.Println("index = ", i, " value = ", num)
	}

	fmt.Println(nums)
	fmt.Println(sum)

	var s string
	s = "hello world"

	for i, c := range s {
		fmt.Println(i, " -> ", c)
	}

}
