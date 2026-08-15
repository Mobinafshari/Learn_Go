package main

import "fmt"

func testDefer() {

	fmt.Println("Start")
	defer fmt.Println("first defer")
	defer fmt.Println("second defer")
	fmt.Println("End")

}

func main() {
	defer fmt.Println("Hello Darkness!")
	defer func() {
		fmt.Println("Hello function!")
	}()
	testDefer()
}
