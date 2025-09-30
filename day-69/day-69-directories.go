package main

import (
	"fmt"
	"os"
)

func main() {
	// create new directory
	err := os.Mkdir("testdir", 0755)
	if err != nil {
		panic(err)
	}
	fmt.Println("Directory created: testdir")

	// create nested directories
	err = os.MkdirAll("parent/child/grandchild", 0755)
	if err != nil {
		panic(err)
	}
	fmt.Println("Nested directories created")

	// remove directory
	err = os.Remove("testdir")
	if err != nil {
		panic(err)
	}
	fmt.Println("Directory removed: testdir")

	// remove nested directories
	err = os.RemoveAll("parent")
	if err != nil {
		panic(err)
	}
	fmt.Println("All nested directories removed")
}
