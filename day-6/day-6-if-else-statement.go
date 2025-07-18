package main

import "fmt"

func main() {

	// Simple if-else statement
	if 7%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// Variables can be declared and initialized directly, then condition can be applied.
	if num := 9; num < 0 {
		fmt.Println("Negative")
	} else if num < 10 {
		fmt.Println("Single-digit")
	} else {
		fmt.Println("Double-digit or more")
	}
}

// Note that you don’t need parentheses around conditions in Go, but that the braces are required.
// There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
