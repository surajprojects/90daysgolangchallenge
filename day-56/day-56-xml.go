package main

import (
	"encoding/xml"
	"fmt"
)

type User struct {
	XMLName xml.Name `xml:"user"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
}

func main() {
	// Struct to XML
	user := User{Name: "Tiger", Age: 22, Email: "tiger@example.com"}
	xmlData, _ := xml.MarshalIndent(user, "", "  ")
	fmt.Println("XML Encode:\n", string(xmlData))

	// XML to Struct
	data := `<user><name>Suraj</name><age>23</age><email>suraj@example.com</email></user>`
	var u User
	xml.Unmarshal([]byte(data), &u)
	fmt.Println("XML Decode:", u)
}
