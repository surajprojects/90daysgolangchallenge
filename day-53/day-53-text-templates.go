package main

import (
	"os"
	"text/template"
)

func main() {
	// Simple template string
	tmpl, _ := template.New("test").Parse("Hello, {{.}}!")
	tmpl.Execute(os.Stdout, "Tiger")

	// Struct data with template
	type User struct {
		Name string
		Age  int
	}

	user := User{Name: "Suraj", Age: 22}
	tmpl2, _ := template.New("userTemplate").Parse("Name: {{.Name}}, Age: {{.Age}}\n")
	tmpl2.Execute(os.Stdout, user)
}
