package main

import "fmt"

type Number interface {
	int | float64 | float32
}

func sum[T Number](numbers ...T) T {
	var res T
	for _, num := range numbers {
		res += num
	}
	return res
}

func main() {

	fmt.Println(sum(30.2, 1.1, 10))
}
