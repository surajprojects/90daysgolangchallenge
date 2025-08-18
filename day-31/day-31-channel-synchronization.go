package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Worker: starting work...")
	time.Sleep(2 * time.Second)
	fmt.Println("Worker: finished work!")
	// notify main goroutine
	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)

	// wait for worker to finish
	<-done
	fmt.Println("Main: all done!")
}
