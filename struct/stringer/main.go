package main

import "fmt"

type Person interface {
	GetName() string
}

type BusinessPerson struct {
	ID   int
	Name string
}

func (e BusinessPerson) String() string {
	return e.Name
}

func main() {

	hashem := BusinessPerson{
		ID:   1,
		Name: "hashem",
	}

	fmt.Println(hashem)

}
