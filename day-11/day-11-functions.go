package main

import "fmt"

// Normal function
func add(a int, b int) int {
	return a + b
}

// Multiple return values
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Add:", add(3, 4))

	a, b := swap("Tiger", "X")
	fmt.Println("Swap:", a, b)

	x, y := split(17)
	fmt.Println("Split:", x, y)
}
