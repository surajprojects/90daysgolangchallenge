package main

import "fmt"

// A function that will take an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Total:", total)
}

func main() {
	sum(1, 2)
	sum(3, 4, 5, 6)
}
