package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone int
}

var contactList []Contact
var contactIndexByName map[string]int
var NextID = 1

func init() {
	contactList := make([]Contact, 0)
	contactIndexByName := make(map[string]int)
}

func addContact(phone int, email string, name string) {
	if _, existed := contactIndexByName[name]; existed {
		fmt.Println("Contact already existed")
		return
	}
	contact := Contact{
		ID:    NextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	contactList = append(contactList, contact)
	contactIndexByName[name] = len(contactList) - 1
	NextID++
}

func findContact(name string) *Contact {
	if index, existed := contactIndexByName[name]; existed {
		return &contactList[index]
	}
	fmt.Println("Contact not founded")
	return nil
}

func main() {
	addContact(9, "22", "dadasd")

}
