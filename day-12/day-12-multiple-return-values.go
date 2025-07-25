package main

import "fmt"

// Function returning two values
func getNameAge() (string, int) {
	return "Tiger", 22
}

func main() {
	name, age := getNameAge()
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}
