package main

import "fmt"

func modifyRef(num *int) {
	if num == nil {
		return
	}
	*num = *num * 10
	fmt.Printf("%+v\n", num)

}

func main() {
	age := 10
	modifyRef(&age)
	fmt.Println(age)

}
