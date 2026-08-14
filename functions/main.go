package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return n
	}
	return n * factorial(n-1)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func nextSeq() int {
	i := 0
	i++
	return i
}

func main() {
	intS := intSeq()
	fmt.Println(intS())
	fmt.Println(intS())
	fmt.Println(intS())
	fmt.Println(intS())
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
}
