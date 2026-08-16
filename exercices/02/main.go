package main

import (
	"fmt"
	"strings"
)

type MathError struct {
	Operation string
	InputA    int
	InputB    int
	Message   string
}

const (
	division    = "Division"
	divisionErr = "Division by zero is not allowed"
)

func (e MathError) Error() string {
	var inputs []string
	if e.Operation == division {
		inputs = append(inputs, fmt.Sprintf("a=%d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("b=%d", e.InputB))
	}
	return fmt.Sprintf("Math error in %s (%s) : %s", e.Operation, strings.Join(inputs, ","), e.Message)
}

func sum(numbers ...int) int {
	total := 0
	defer fmt.Printf("sum finished ")
	for _, n := range numbers {
		total += n
	}
	fmt.Println(total)
	return total
}

func divisionOperation(a, b int) (int, error) {
	if b == 0 {
		return 0, &MathError{
			Operation: division,
			InputA:    a,
			InputB:    b,
			Message:   divisionErr,
		}
	}
	return a / b, nil
}

func main() {
	sum(10, 2)

	value, err := divisionOperation(10, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}
