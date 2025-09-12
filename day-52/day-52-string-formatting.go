package main

import (
	"fmt"
)

func main() {
	// Normal string print
	name := "Tiger"
	age := 22

	// Using Printf with format specifiers
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Formatting numbers
	pi := 3.141592
	fmt.Printf("Pi up to 2 decimals: %.2f\n", pi)

	// Formatting with width
	num := 42
	fmt.Printf("Number with padding: %5d\n", num)

	// Boolean formatting
	isCool := true
	fmt.Printf("Am I cool? %t\n", isCool)

	// Using Sprintf to return formatted string
	message := fmt.Sprintf("Hello, %s 👋", name)
	fmt.Println(message)
}
