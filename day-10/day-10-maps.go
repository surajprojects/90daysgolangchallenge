package main

import "fmt"

func main() {
	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax.
	m["Tiger"] = 100
	m["Insights"] = 88

	// Printing a map with e.g. fmt.Println will show all of its key/value pairs.
	fmt.Println("map:", m)

	// Get a value for a key with name[key].
	v1 := m["Tiger"]
	fmt.Println("Tiger score:", v1)

	// The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len:", len(m))

	// The builtin delete removes key/value pairs from a map.
	delete(m, "Insights")
	fmt.Println("map after delete:", m)

	// To remove all key/value pairs from a map, use the clear builtin.
	clear(m)
	fmt.Println("map:", m)

	// The optional second return value when getting a value from a map indicates if the key was present in the map. Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	_, present := m["Insights"]
	fmt.Println("Insights present?", present)

	// You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
