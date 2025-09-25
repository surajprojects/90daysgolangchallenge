package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "TigerX"

	// Encode to Base64
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded:", encoded)

	// Decode back to string
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println("Decoded:", string(decoded))
}
