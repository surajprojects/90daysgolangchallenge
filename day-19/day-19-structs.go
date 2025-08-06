package main

import "fmt"

// Define a struct
type Person struct {
	name string
	age  int
}

func main() {
	// Create a struct instance
	p1 := Person{name: "Tiger", age: 22}
	fmt.Println("Person 1:", p1)

	// Access fields
	fmt.Println("Name:", p1.name)
	fmt.Println("Age:", p1.age)

	// Modify field
	p1.age += 1
	fmt.Println("Updated Age:", p1.age)
}
