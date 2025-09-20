package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// Formatting
	fmt.Println("Default format:", t.String())
	fmt.Println("RFC3339:", t.Format(time.RFC3339))

	// Custom layout
	fmt.Println("12-hour clock:", t.Format("03:04PM"))
	fmt.Println("Date only:", t.Format("2006-01-02"))
	fmt.Println("Day/Month/Year:", t.Format("02-Jan-2006"))

	// Parsing
	timeStr := "2025-08-18T10:15:00+05:30"
	parsed, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Parsed time:", parsed)
}
