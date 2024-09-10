package main

import "fmt"

type car struct{
make string
 model string
 year int
 engine string
 horsepower int
 color string
 price float64 
}

func (f car) favourite() string {
return "My favourite car is " + f.make
} 

func main() {
	Car := car{
		make: "Toyota",
		model: "landcruiser",
		year: 2017,
		engine: "v8",
		horsepower: 500,
		color: "black",
		price: 2500000 ,
	}
	fmt.Println(Car.favourite())
}