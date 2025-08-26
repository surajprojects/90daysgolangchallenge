package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a ticker that ticks every 1 second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Run for 5 seconds only
	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		case <-done:
			fmt.Println("Ticker stopped")
			return
		}
	}
}
