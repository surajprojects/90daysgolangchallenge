package main

import "fmt"

func main() {
	// Create a channel of type string
	messages := make(chan string)

	// Send a value into the channel in a goroutine
	go func() {
		messages <- "Hello from Goroutine!"
	}()

	// Receive the value from the channel
	msg := <-messages
	fmt.Println("Received:", msg)
}
