package main

import "fmt"

func main() {
	ch := make(chan string)

	// Non-blocking receive
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received")
	}

	// Non-blocking send
	select {
	case ch <- "hello":
		fmt.Println("Message sent")
	default:
		fmt.Println("No receiver ready")
	}
}
