package main

import (
	"fmt"
	"os"
)

func main() {
	// Get environment variable
	user := os.Getenv("USER")
	fmt.Println("Current user:", user)

	// Set environment variable (only works during runtime of this program)
	os.Setenv("APP_MODE", "production")
	fmt.Println("App Mode:", os.Getenv("APP_MODE"))

	// Check if variable exists
	secret, present := os.LookupEnv("SECRET_KEY")
	if present {
		fmt.Println("Secret Key Found:", secret)
	} else {
		fmt.Println("No Secret Key set")
	}
}
