package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current time:", now)

	// Create a specific time
	then := time.Date(2002, 8, 23, 5, 30, 0, 0, time.UTC)
	fmt.Println("My Birthday:", then)

	// Extract fields
	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())

	// Comparison
	fmt.Println("Before?", then.Before(now))
	fmt.Println("After?", then.After(now))
	fmt.Println("Equal?", then.Equal(now))

	// Difference
	diff := now.Sub(then)
	fmt.Println("Since Birthday:", diff.Hours(), "hours")

	// Add/Subtract
	fmt.Println("1 Month Later:", now.AddDate(0, 1, 0))
}
