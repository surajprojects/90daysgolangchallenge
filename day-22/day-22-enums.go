package main

import "fmt"

// Enum-like constants using iota
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// String method to print day names
func (d Weekday) String() string {
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[d]
}

func main() {
	today := Tuesday
	fmt.Println("Today is:", today)
}
