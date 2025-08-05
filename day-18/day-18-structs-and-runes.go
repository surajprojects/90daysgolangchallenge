package main

import "fmt"

func main() {
	str := "TigerX🔥"

	fmt.Println("Original string:", str)
	fmt.Println("Length in bytes:", len(str)) // byte length

	fmt.Println("Runes:")
	for i, r := range str {
		fmt.Printf("Index %d: Rune %c Unicode %U\n", i, r, r)
	}
}
