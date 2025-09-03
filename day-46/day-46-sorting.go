package main

import (
	"fmt"
	"sort"
)

func main() {
	// Integers
	nums := []int{7, 2, 9, 4, 1}
	sort.Ints(nums)
	fmt.Println("Sorted ints:", nums)

	// Strings
	words := []string{"banana", "apple", "cherry"}
	sort.Strings(words)
	fmt.Println("Sorted strings:", words)

	// Check if sorted
	fmt.Println("Is nums sorted?", sort.IntsAreSorted(nums))
}
