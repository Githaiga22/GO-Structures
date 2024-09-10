# GO Structures

### 1. Introduction to Go Structs

- What is a Struct?

  - In Go, a struct is a collection of fields or variables, which can be of different types, that are grouped together to represent a more complex entity.
  - Think of a struct as a blueprint or template used to create a more complex data structure like a Person, Car, or Book.

- Why Use Structs?

   - To organise and bundle related data together.
   - They allow you to model real-world things like a customer or product, each having multiple attributes (fields).
### 2. Declaring a Struct
**Syntax:**
```bash
type StructName struct {
field1 DataType
field2 DataType
...
}
```
Example 1: A simple Person struct
```bash
type Person struct { 
Name string
 Age int
 Gender string 
}
```
This example defines a Person struct with three fields: Name (string), Age (int), and Gender (string).


### 3. Creating and Initializing a Struct

**Method 1: Using the Field Names**
   - You can create an instance of a struct by explicitly specifying the field names and values.

**Example:**
```bash
person := Person{
	Name:   "John Doe",
	Age:	25,
	Gender: "Male",
}
```
testing
```bash
package main

import "fmt"

type Person struct {
	name string
	age int
	gender string
}

func main() {
	person := Person{
		name : "Allan kamau",
		age : 23,
		gender : "male",

	}
fmt.Println("Name:", person.name)
fmt.Println("age:", person.age)
fmt.Println("gender:", person.gender)
}

//output
Name: Allan kamau
age: 23
gender: male
```

### 4. Accessing Struct Fields

Once a struct is created, you can access its fields using the dot (.) operator.

**Example:**
```bash
fmt.Println("Name:", person.name)
fmt.Println("age:", person.age)
fmt.Println("gender:", person.gender)
```
### 5. Modifying Struct Fields

   - You can also modify the fields of a struct after it has been created.

**Example 1:**
   ```bash
   package main

import "fmt"

type Person struct {
	name string
	age int
	gender string
}

func main() {
	person := Person{
		name : "Allan kamau",
		age : 23,
		gender : "male",

	}
fmt.Println("Name:", person.name)
fmt.Println("age:", person.age)
fmt.Println("gender:", person.gender)

person.age = 26
//modifying agefields
fmt.Println("updated Age:", person.age)

//output
Name: Allan kamau
age: 23
gender: male
updated Age: 26
}
```
### 6. Struct Methods

Go allows you to define methods on structs, just like functions, but these methods are tied to a struct and can access its fields.

**Example 2:**

 Adding a Method to Person
 ```bash
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

//output
Hello, my name is Allan kamau
}
```
### 7. Struct with Nested Structs

Sometimes, structs can contain other structs as fields, allowing for more complex data modeling.

**Example 3:**

A Student struct containing an Address struct
```bash
package main

import "fmt"

type Address struct {
	city string
	state string
}
type Student struct {
	Name string
	Age int
	Address Address
}
func main() {
	student := Student{
		Name: "Fred Gitonga",
		Age: 21,
		Address: Address{
			city: "Kisumu",
			state: "Mamboleo",
		},
	}
fmt.Println(student.Name, "lives in", student.Address.city)
}

//ouput
fred gitonga lives in kisumu
```
### 8. Structs with Functions as Fields

You can even add functions to a struct as fields, allowing for flexible behaviors inside your structs.

**Example 4:**

 Calculator struct with function fields
 ```bash
 package main

import "fmt"


type Calculator struct {
	Add func(a, b int)  int
}

func main() {
	calculator  := Calculator{
		Add: func(a, b int) int {
			return a + b
		},
	}
result := calculator.Add(3, 5)
fmt.Println("Sum:", result)
}

//output
Sum: 8
```
### 9. Pointers to Structs

In Go, structs can be passed by value or by reference (using pointers). Passing by reference allows you to modify the original struct.

**Example 5:** 

Passing a Pointer to a Struct
```bash
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

//output
26
```

### 10. Common Mistakes with Structs

- Uninitialized Fields: If you don't initialize a field, Go automatically assigns a zero value (empty string, 0, false, etc.).
- Using pointers when not needed: For small structs, passing by value may be more efficient.
   

### 11. Practice Problems

- Easy: Create a Book struct with fields Title, Author, and Pages. Write a method to display the book's info.
- Intermediate: Create a Library struct that holds multiple Book structs. Write methods to add a book and list all books in the library.
- Hard: Create a Vehicle struct with fields like Make, Model, and Year. Then create a Fleet struct that manages multiple vehicles, including adding, removing, and searching vehicles by year.

### 12. Conclusion

Key Points to Remember:
  - Structs are used to group related data.
  - Use methods to add functionality to structs.
 - Pointers allow you to modify structs in place.
  - Structs can be nested or contain functions as fields.

  ## Author
  Allan Githaiga Kamau

  ### Github profile
  [Githaiga22](https://github.com/Githaiga22)