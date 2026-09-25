package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" {
		return "No Address Provided"
	}
	return fmt.Sprintf("%s", "%s", "%s", "%s", a.City, a.State, a.Street, a.ZipCode)
}

type Customer struct {
	CustomerId      int
	Name            string
	Email           string
	BillingAddress  Address
	ShippingAddress Address
}

func (c Customer) PrintDetails() {
	fmt.Printf("CustomerId: ", c.CustomerId)
	fmt.Printf("Name: ", c.Name)
	fmt.Printf("Email: ", c.Email)
	fmt.Printf("Billing Address: ", c.BillingAddress.FullAddress())
	fmt.Printf("Shipping Address: ", c.ShippingAddress.FullAddress())
}

func main() {

	cus1 := Customer{
		CustomerId: 1,
		Name:       "hashem",
		Email:      "",
		BillingAddress: Address{
			Street:  "",
			City:    "hashem abad",
			State:   "",
			ZipCode: "",
		},
		ShippingAddress: Address{
			Street:  "",
			City:    "hashem abad",
			State:   "",
			ZipCode: "",
		},
	}

	cus1.PrintDetails()
}
