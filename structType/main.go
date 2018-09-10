package main

import (
	"fmt"
)

type contact struct {
	phone string
	email string
}

type person struct {
	firstName   string
	lastName    string
	contactInfo contact
}

func main() {

	jim := person{
		firstName: "shyam",
		lastName:  "singh",
		contactInfo: contact{
			phone: "87225",
			email: "spdinf",
		},
	}

	// Shortcut for pointer, go automatically pass the memory address to *person in updateNamUsingPointer function.
	jim.updateNamUsingPointer("ShyamShortcut")

	// one way to use pointer
	// jimPointer := &jim
	// jimPointer.updateNamUsingPointer("NewPointerName")

	// res := updateNormally(jim, "ShyamNew")
	// fmt.Printf("%+v", res)

	// jim.firstName = "JimmyNew"

	// jim.updatePerson("Jimmy")

	jim.print()
	// fmt.Println(jim.contactInfo.phone)

	// fmt.Printf("%+v", jim)

	// // shyam := person{firstName: "Shyam", lastName: "singh"}
	// // patientName := shyam.firstName

	// // another way to initialise struct
	// var shyam person

	// shyam.firstName = "Shyam"
	// shyam.lastName = "Singh"

	// fmt.Println(shyam)
}

// up

// creating a receiver for struct type to print all details

func (p person) print() {
	fmt.Printf("%+v", p)
}

// updating through normal way, it has happend, problem only in receiver function
func updateNormally(per person, fName string) person {
	per.firstName = fName
	return per
}

// update using poniter

func (p *person) updateNamUsingPointer(firstName string) {
	(*p).firstName = firstName

}

// updating value through receiver, it wont simply changed the value, due to receiver implementation

func (p person) updatePerson(firstNewName string) {
	p.firstName = firstNewName

}
