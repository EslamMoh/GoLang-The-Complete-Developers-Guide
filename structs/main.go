package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// Definition of struct of type person
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}

	// jimPointer := &jim // Give us access to the memory address that jim variable is pointing at
	jim.updateName("jimmy") // We called the function through the root person type
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName // Whenever we use * before a pointer, this turns the pointer into value
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
