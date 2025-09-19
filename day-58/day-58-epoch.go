package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Now:", now)

	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println("Seconds since Epoch:", secs)
	fmt.Println("Nanoseconds since Epoch:", nanos)

	// Convert back
	millis := nanos / 1e6
	fmt.Println("Milliseconds since Epoch:", millis)

	// From epoch back to Time
	t := time.Unix(secs, 0)
	fmt.Println("Restored time:", t)
}
