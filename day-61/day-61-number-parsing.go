package main

import (
	"fmt"
	"strconv"
)

func main() {
	// String to int
	numStr := "42"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Conversion error:", err)
	} else {
		fmt.Println("String to int:", num)
	}

	// String to float
	floatStr := "3.14"
	f, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Conversion error:", err)
	} else {
		fmt.Println("String to float:", f)
	}

	// Invalid parse example
	badStr := "hello"
	_, err = strconv.Atoi(badStr)
	if err != nil {
		fmt.Println("Invalid parse:", err)
	}
}
