package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://example.com:8080/path?user=Tiger&age=22#section"

	// Parse URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	// Access parts
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Fragment:", parsedURL.Fragment)

	// Query parameters
	q := parsedURL.Query()
	fmt.Println("User:", q.Get("user"))
	fmt.Println("Age:", q.Get("age"))
}
