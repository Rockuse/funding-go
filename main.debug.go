package main

import "fmt"

type Address struct {
	Street  string
	City    string
	ZipCode string
}

func (a *Address) PrintAddress() {
	fmt.Printf("Address: %s, %s, %s\n", a.Street, a.City, a.ZipCode)
}

type Person struct {
	FirstName string
	LastName  string
	Address   // Embedded struct
}

func (p *Person) PrintAddress() {
	fmt.Printf("Person: %s, %s\n", p.FirstName, p.LastName)
}

func Tester() {
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "10001",
		},
	}

	person.Address.PrintAddress() // Call the method of the embedded struct
	person.PrintAddress()         // Call the method of the Person struct
}
