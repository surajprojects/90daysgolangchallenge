package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := "TigerXInsights"

	// Create SHA256 hash
	hash := sha256.Sum256([]byte(data))

	// Print in hex format
	fmt.Printf("Data: %s\n", data)
	fmt.Printf("SHA256: %x\n", hash)
}
