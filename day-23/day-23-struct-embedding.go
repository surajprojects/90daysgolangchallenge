package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person     // Embedded struct
	EmployeeID string
}

func main() {
	emp := Employee{
		Person:     Person{Name: "Tiger", Age: 22},
		EmployeeID: "EMP001",
	}

	fmt.Println("Name:", emp.Name) // Accessed directly!
	fmt.Println("Age:", emp.Age)
	fmt.Println("Employee ID:", emp.EmployeeID)
}
