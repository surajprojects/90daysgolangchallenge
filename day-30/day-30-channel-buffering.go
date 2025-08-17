package main

import "fmt"

func main() {
	// Create a buffered channel with capacity 2
	messages := make(chan string, 2)

	// Send values into the buffered channel
	messages <- "Tiger"
	messages <- "X"

	// Receive values from the channel
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
