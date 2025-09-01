package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		counter int
		mu      sync.Mutex
		wg      sync.WaitGroup
	)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()   // Lock karna zaroori hai
				counter++   // Shared resource ko safe tarike se update karte hain
				mu.Unlock() // Unlock karna mat bhoolna
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
