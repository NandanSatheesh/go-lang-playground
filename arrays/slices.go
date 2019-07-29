package main

import "fmt"

func main() {

	s := make([]string, 6)

	s[0] = "A"
	s[1] = "B"
	s[2] = "C"
	s[3] = "D"
	s[4] = "E"
	s[5] = "F"

	s = append(s, "G", "H")

	fmt.Println(s)

	a := make([]string, 3)
	copy(a, s[:3])

	fmt.Println(s[:4])

	fmt.Println(a)

	twoD := make([][]int, 5)

	fmt.Println(twoD)

	for i := 0; i < 5; i++ {
		twoD[i] = make([]int, i)
		for j := 0; j < i; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

}
