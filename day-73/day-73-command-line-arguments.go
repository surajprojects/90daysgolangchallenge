package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args gives us command-line arguments as a slice
	args := os.Args

	// First element is always the program name
	fmt.Println("Program name:", args[0])

	// Print all arguments
	fmt.Println("Arguments:", args[1:])

	// Access individual arguments
	if len(args) > 1 {
		fmt.Println("First argument:", args[1])
	}
}
