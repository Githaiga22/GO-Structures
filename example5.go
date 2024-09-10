package main 

import "fmt"

type Person struct {
	Name string
	Age int
	gender string
}

func UpdateAge(p *Person, newAge int) {
	p.Age = newAge
}
func main() {
	john := Person{ Name: "john", Age: 25, gender: "Male"}
	UpdateAge(&john, 26)
	fmt.Println(john.Age) 

}