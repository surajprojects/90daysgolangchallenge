package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Basic random numbers
	fmt.Println("Random int:", rand.Intn(100)) // 0–99
	fmt.Println("Random float:", rand.Float64())

	// Random range scaling
	fmt.Println("Float 0–5:", rand.Float64()*5)

	// Using a seed for different results each run
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Seeded int:", rand.Intn(100))
	fmt.Println("Seeded float:", rand.Float64())
}
