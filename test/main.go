package main

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

var users []User

func createUser(name, email string, age int) {
	user := User{
		Name:  name,
		Email: email,
		Age:   age,
		ID:    len(users),
	}
	users = append(users, user)
}

func updateUser(u *User) {
	u.Name = "Random"

}

func deleteUser(id int) {

	for _, user := range users {
		if user.ID == id {

		}
	}
}

func main() {

}
