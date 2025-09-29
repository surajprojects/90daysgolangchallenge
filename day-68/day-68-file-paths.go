package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	p := filepath.Join("dir1", "subdir1", "file.txt")
	fmt.Println("Path:", p)

	fmt.Println("Base:", filepath.Base(p))
	fmt.Println("Dir:", filepath.Dir(p))
	fmt.Println("IsAbs:", filepath.IsAbs(p))

	abs, err := filepath.Abs(p)
	if err != nil {
		panic(err)
	}
	fmt.Println("Absolute:", abs)

	ext := filepath.Ext(p)
	fmt.Println("Ext:", ext)
}
