package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 1)

	// Simulate a goroutine that takes some time
	go func() {
		time.Sleep(2 * time.Second)
		c <- "Result mil gaya!"
	}()

	select {
	case res := <-c:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout ho gaya ⏳")
	}
}
