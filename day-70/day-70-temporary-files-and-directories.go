package main

import (
	"fmt"
	"os"
)

func main() {
	// Temporary file
	tmpFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name()) // cleanup
	fmt.Println("Temporary file created:", tmpFile.Name())

	// Temporary directory
	tmpDir, err := os.MkdirTemp("", "example-dir-*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpDir) // cleanup
	fmt.Println("Temporary directory created:", tmpDir)
}
