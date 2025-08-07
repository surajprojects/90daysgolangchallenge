package main

import "fmt"

// Define struct
type Rectangle struct {
	width  float64
	height float64
}

// Method with value receiver
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Method with pointer receiver
func (r *Rectangle) scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func main() {
	r := Rectangle{width: 5, height: 3}
	fmt.Println("Original Rectangle:", r)
	fmt.Println("Area:", r.area())

	r.scale(2)
	fmt.Println("Scaled Rectangle:", r)
	fmt.Println("New Area:", r.area())
}
