package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("This will print")

	// Exit program with status code 0 (success)
	os.Exit(0)

	// Ye line print nahi hogi
	fmt.Println("This will NOT print")
}
