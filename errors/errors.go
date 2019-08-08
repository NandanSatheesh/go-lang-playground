package main

import (
	"errors"
	"fmt"
)

func throwError(input int) (int, error) {

	if input != 23 {
		return 0, errors.New("This is wrong. How can you do this?")
	}

	return 0, nil
}

func main() {

	_, err := throwError(123)
	fmt.Println(err)

}
