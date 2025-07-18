package main

import "fmt"

func main() {

	// While style for loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// Classic for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Infinite for loop, use break keyword to exit
	for {
		fmt.Println("loop")
		break
	}

	// Range for loop
	for i := range 3 {
		fmt.Println("range", i)
	}
}
