package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// func main() {
// 	var alex person

// 	alex.firstName = "Alex"
// 	alex.lastName = "Anderson" // dot notation to assign values
// 	// alex := person{firstName: "Alex", lastName: "Anderson"} to create new struct person alex
// 	fmt.Println(alex) // prints the struct
// 	// fmt.Printf("%+v", alex) shows values from struct.
// }

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Parsons",
		contact: contactInfo{
			email:   "jim@parsons.com",
			zipCode: 89000,
		},
	}
	// jimPointer := &jim  getting the memory address of the struct jim (pointer)
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) { // *pointer (type) type description, working with a pointer to a person
	(*pointerToPerson).firstName = newFirstName // turning memory address back into a value (pointer)
}

func (p person) print() {
	fmt.Println(p)
	fmt.Printf("%+v", p)
}

// reference types dont need pointers.  (slices, maps, channels, pointers, functions)

// value types need pointers to change (int, float, string, bool, structs)
