package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	daniel := person{
		firstName: "Daniel",
		lastName:  "Santos",
		contactInfo: contactInfo{
			email:   "drsantos20@gmail.com",
			zipCode: 94000,
		},
	}

	daniel.updateName("dani")
	daniel.print()
}

// & is a memory copy
// * is a value access
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
