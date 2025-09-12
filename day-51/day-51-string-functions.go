package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("golang", "go"))            // true
	fmt.Println(strings.HasPrefix("golang", "go"))           // true
	fmt.Println(strings.HasSuffix("golang", "lang"))         // true
	fmt.Println(strings.Count("golang", "g"))                // 2
	fmt.Println(strings.Index("golang", "o"))                // 1
	fmt.Println(strings.ToUpper("golang"))                   // GOLANG
	fmt.Println(strings.ToLower("GoLang"))                   // golang
	fmt.Println(strings.ReplaceAll("go go go", "go", "run")) // run run run
	fmt.Println(strings.Split("a,b,c", ","))                 // [a b c]
}
