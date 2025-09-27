package main

import (
	"fmt"
	"os"
)

func main() {
	// File content read karna
	data, err := os.ReadFile("day-65/example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print file ka content
	fmt.Println(string(data))
}
