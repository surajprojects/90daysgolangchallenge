package main

import "fmt"

func mayPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("Something went wrong!")
}

func main() {
	mayPanic()
	fmt.Println("Program continues after recover 🚀")
}
