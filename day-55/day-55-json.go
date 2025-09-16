package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// Struct to JSON
	user := User{Name: "Tiger", Age: 22, Email: "tiger@example.com"}
	jsonData, _ := json.Marshal(user)
	fmt.Println("JSON Encode:", string(jsonData))

	// JSON to Struct
	data := `{"name":"Suraj","age":23,"email":"suraj@example.com"}`
	var u User
	json.Unmarshal([]byte(data), &u)
	fmt.Println("JSON Decode:", u)
}
