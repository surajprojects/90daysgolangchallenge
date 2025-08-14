package main

import (
	"fmt"
)

// Custom error type
type DivideError struct {
	Dividend float64
	Divisor  float64
	Message  string
}

// Implement the Error() method
func (e *DivideError) Error() string {
	return fmt.Sprintf("Error: %s | Dividend: %.2f, Divisor: %.2f", e.Message, e.Dividend, e.Divisor)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &DivideError{
			Dividend: a,
			Divisor:  b,
			Message:  "cannot divide by zero",
		}
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(5, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result:", result)
	}
}
