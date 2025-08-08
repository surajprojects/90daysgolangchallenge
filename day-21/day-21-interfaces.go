package main

import "fmt"

// Interface
type Shape interface {
	area() float64
}

// Structs
type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

// Method implementations
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Function that accepts interfaceGenerics
func printArea(s Shape) {
	fmt.Println("Area:", s.area())
}

func main() {
	c := Circle{radius: 5}
	r := Rectangle{width: 4, height: 3}

	printArea(c)
	printArea(r)
}
