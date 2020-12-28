package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact contactInfo
	contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// var alex person
	// // fmt.Println(alex)

	// alex.firstName = "mk"
	// alex.lastName = "thfsd"

	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jom",
		lastName:  "sth",
		contactInfo: contactInfo{
			email:   "jimipur@gmail.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("jimyy")

	jim.updateName("jimyy")

	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
