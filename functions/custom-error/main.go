package main

import (
	"errors"
	"fmt"
)

var invalidA = errors.New("a is too large")
var zeroB = errors.New("b cannot be zero")

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, zeroB
	}

	if a > 1000 {
		return 0, invalidA
	}
	return a / b, nil

}

func main() {

	res, error := divide(1001, 30)
	if error != nil {
		if errors.Is(error, invalidA) {
			fmt.Println("enter another a pls")
		}
	}
	fmt.Println(res)

}
