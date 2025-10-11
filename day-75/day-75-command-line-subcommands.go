package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define subcommands
	helloCmd := flag.NewFlagSet("hello", flag.ExitOnError)
	helloName := helloCmd.String("name", "Guest", "Your name")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addX := addCmd.Int("x", 0, "First number")
	addY := addCmd.Int("y", 0, "Second number")

	// Check which subcommand is called
	if len(os.Args) < 2 {
		fmt.Println("expected 'hello' or 'add' subcommands")
		return
	}

	switch os.Args[1] {
	case "hello":
		helloCmd.Parse(os.Args[2:])
		fmt.Printf("Hello, %s!\n", *helloName)

	case "add":
		addCmd.Parse(os.Args[2:])
		fmt.Printf("Sum: %d\n", *addX+*addY)

	default:
		fmt.Println("Unknown command")
	}
}
