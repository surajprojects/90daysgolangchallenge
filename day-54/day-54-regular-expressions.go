package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Compile regex pattern
	re := regexp.MustCompile("[a-zA-Z]+")

	// Match string
	fmt.Println("Match:", re.MatchString("TigerX123"))

	// Find first match
	fmt.Println("Find:", re.FindString("TigerX123"))

	// Find all matches
	fmt.Println("FindAll:", re.FindAllString("Go123Lang456Tiger", -1))

	// Replace matches
	fmt.Println("Replace:", re.ReplaceAllString("123GoLang456", "X"))
}
