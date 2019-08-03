package main

import "fmt"

func main() {
	x := 20
	withoutPointers(x)
	withPointer(&x)
}

func withoutPointers(ival int) {
	ival = 10
	fmt.Println(ival)
}

func withPointer(ival *int) {
	*ival++
	fmt.Println(ival)
	fmt.Println(*ival)
}
