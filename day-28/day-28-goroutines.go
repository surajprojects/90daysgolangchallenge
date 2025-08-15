package main

import (
	"fmt"
	"time"
)

func printMsg(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// Normal function call
	printMsg("Synchronous Call")

	// Goroutine
	go printMsg("Goroutine Call")

	// Another Goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("Inline Goroutine")

	// Give goroutines time to run
	time.Sleep(time.Second * 2)
	fmt.Println("Main function done!")
}
