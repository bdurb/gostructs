package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson" // dot notation to assign values
	// alex := person{firstName: "Alex", lastName: "Anderson"} to create new struct person alex
	fmt.Println(alex) // prints the struct
	// fmt.Printf("%+v", alex) shows values from struct.
}
