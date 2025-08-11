package main

import "fmt"

// Generic function with type parameter
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	intSlice := []int{1, 2, 3}
	stringSlice := []string{"Tiger", "GoLang", "Generics"}

	fmt.Println("Integer slice:")
	PrintSlice(intSlice)

	fmt.Println("String slice:")
	PrintSlice(stringSlice)
}
