package main

import "fmt"

// This fact function calls itself until it reaches the base case of fact(0).
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println("Factorial of 5:", fact(5)) // Output: 120

	// Anonymous functions can also be recursive, but this requires explicitly declaring a variable with var to store the function before it’s defined.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		// Since fib was previously declared in main, Go knows which function to call with fib here.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println("Fibonacci of 7:", fib(7)) // Output: 13
}
