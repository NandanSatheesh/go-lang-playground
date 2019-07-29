package main

import "fmt"

func main() {

	s := make(map[string]int)

	s["a"] = 1
	s["b"] = 2

	fmt.Println(s)

	delete(s, "a")
	fmt.Println(s)

	_, p := s["123"]
	n := map[string]int{"a": 123, "b": 234}
	fmt.Println(n)
	fmt.Println(p)
}
