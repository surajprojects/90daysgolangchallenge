package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("Hello Tiger 🐯, GoLang writing files is easy!\n")

	// Write file (overwrite mode)
	err := os.WriteFile("example.txt", data, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File written successfully!")
}
