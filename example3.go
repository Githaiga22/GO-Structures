package main

import "fmt"

type Person struct {
	name string
	age int
	gender string
	Address Address
}
type Address struct {
	location string
	zipcode int
}

func main() {
	person := Person{
		name : "Allan kamau",
		age : 23,
		gender : "male",
		Address: Address{
			location: "kisumu",
			zipcode: 40100,
		},

		

	}


	fmt.Println(person.name, "lives in", person.Address.location) 
}
//Allan kamau lives in kisumu

