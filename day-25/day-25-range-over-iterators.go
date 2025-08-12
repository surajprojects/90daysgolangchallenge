package main

import "iter" // Go 1.23+ experimental package

// Custom iterator returning numbers 1 to n
func Numbers(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	for num := range Numbers(5) {
		println(num)
	}
}
