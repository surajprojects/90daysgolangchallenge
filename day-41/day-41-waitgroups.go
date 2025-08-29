package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // mark this goroutine as done
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // simulate work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // tell WaitGroup we are adding one goroutine
		go worker(i, &wg)
	}

	wg.Wait() // wait until all goroutines finish
	fmt.Println("All workers finished 🚀")
}
