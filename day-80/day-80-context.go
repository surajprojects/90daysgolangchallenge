package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Run a goroutine that simulates work
	go func() {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Task finished")
		case <-ctx.Done():
			fmt.Println("Context cancelled:", ctx.Err())
		}
	}()

	// Wait for goroutine
	time.Sleep(3 * time.Second)
}
