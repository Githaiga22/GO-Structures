package main

import "fmt"

type Person struct {
	name string
	age int
	gender string
}
//define method greet on person struct
func (p Person) Greet() string {
	return "Hello, my name is " + p.name
}
func main() {
	person := Person{
		name : "Allan kamau",
		age : 23,
		gender : "male",

	}
fmt.Println(person.Greet())
// fmt.Println("age:", person.age)
// fmt.Println("gender:", person.gender)

// person.age = 26
// //modifying agefields
// fmt.Println("updated Age:", person.age)
}

