package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	// Goroutine 2
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	// Using select
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
