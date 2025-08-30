package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple ticker-based rate limiting
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(500 * time.Millisecond) // 1 request per 500ms

	for req := range requests {
		<-limiter
		fmt.Println("Handling request", req, "at", time.Now())
	}
}
