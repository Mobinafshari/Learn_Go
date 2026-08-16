package main

import (
	"fmt"
	"strings"
	"time"
)

type Employee struct {
	ID        int64
	FirstName string
	LastName  string
	Position  string
	IsActive  bool
	Salary    int
	JoinedAt  time.Time
}

func newEmployee(id int64, firstName, lastName, position string, salary int) Employee {
	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		Salary:    salary,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}
}

func (e Employee) fullName() string {
	names := [2]string{e.FirstName, e.LastName}
	return strings.Join(names[:], " ")
}

func (e Employee) deActivate() {
	e.IsActive = false
}
func main() {
	hashem := newEmployee(12, "Hashem", "Naseri ", "Software", 111111)

	fmt.Println(hashem.fullName())
}
