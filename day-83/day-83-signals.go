package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create channel to receive OS signals
	sigs := make(chan os.Signal, 1)

	// Notify channel for specific signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Waiting for signal (Press Ctrl+C)...")

	// Wait for signal in a goroutine
	go func() {
		sig := <-sigs
		fmt.Println("\nReceived signal:", sig)
		fmt.Println("Exiting gracefully 🐯")
		os.Exit(0)
	}()

	// Simulate work
	for i := 1; i <= 5; i++ {
		fmt.Println("Working...", i)
		time.Sleep(1 * time.Second)
	}
}
