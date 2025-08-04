package main

import "fmt"

func main() {
	a := 10
	b := &a // pointer to 'a'

	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", b)
	fmt.Println("Value from pointer b:", *b) // dereferencing

	*b = 20 // updating value through pointer
	fmt.Println("New value of a:", a)
}
