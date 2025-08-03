package main

import "fmt"

func main() {
	// Array
	nums := [3]int{10, 20, 30}
	for i, v := range nums {
		fmt.Println("Index:", i, "Value:", v)
	}

	// Slice
	fruits := []string{"apple", "banana", "mango"}
	for i, fruit := range fruits {
		fmt.Println("Fruit", i, "is", fruit)
	}

	// Map
	scores := map[string]int{"Tiger": 99, "X": 100}
	for key, value := range scores {
		fmt.Println(key, "scored", value)
	}

	// String
	word := "Tiger"
	for i, char := range word {
		fmt.Printf("Character at %d is %c\n", i, char)
	}
}
