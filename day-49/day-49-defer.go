package main

import "fmt"

func main() {
	fmt.Println("Start of main")

	// Ye defer statement tab run hoga jab function exit hoga
	defer fmt.Println("This is deferred!")

	fmt.Println("End of main")
}
