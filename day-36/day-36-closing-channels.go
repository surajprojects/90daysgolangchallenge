package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	// Sending values
	ch <- 1
	ch <- 2
	ch <- 3

	// Close the channel
	close(ch)

	// Range over channel after closing
	for val := range ch {
		fmt.Println(val)
	}
}
