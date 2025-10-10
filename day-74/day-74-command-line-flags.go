package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	name := flag.String("name", "Guest", "Your name")
	age := flag.Int("age", 18, "Your age")
	isCool := flag.Bool("cool", false, "Are you cool?")

	// Parse the flags
	flag.Parse()

	// Use them
	fmt.Println("Name:", *name)
	fmt.Println("Age:", *age)
	fmt.Println("Cool:", *isCool)
}
