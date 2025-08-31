package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Thread-safe increment
			atomic.AddUint64(&ops, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Ops:", ops)
}
