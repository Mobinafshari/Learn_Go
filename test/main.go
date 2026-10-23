package main

import (
	"fmt"
	"slices"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}
type UpdateUserInput struct {
	Name  *string
	Email *string
	Age   *int
}

var users []User
var nextID = 1

func createUser(name, email string, age int) {
	user := User{
		Name:  name,
		Email: email,
		Age:   age,
		ID:    nextID,
	}
	users = append(users, user)
	nextID++
}

func updateUser(id int, input UpdateUserInput) error {
	for i := range users {
		if users[i].ID == id {
			if input.Name != nil {
				users[i].Name = *input.Name
			}

			if input.Email != nil {
				users[i].Email = *input.Email
			}

			if input.Age != nil {
				users[i].Age = *input.Age
			}

			return nil
		}
	}

	return fmt.Errorf("user with ID %d not found", id)
}

func deleteUser(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = slices.Delete(users, i, i+1)
			return nil
		}
	}

	return fmt.Errorf("user with ID %d not found", id)
}

func main() {
	createUser("Hashem", "Hashem@gmail.com", 22)
	fmt.Println(users)
	newName := "Naser"
	err := updateUser(1, UpdateUserInput{
		Name: &newName,
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
	deleteUser(1)
	fmt.Println(users)

}
