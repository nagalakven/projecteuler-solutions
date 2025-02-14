package main

import "fmt"

func main() {
	n := 1001
	sum := 1 //start with 1

	for i := 3; i <= n; i = i + 2 {
		result := (4 * i * i) - (6 * (i - 1))
		sum += result
	}

	fmt.Printf("Number Spiral Diagonals - Sum of diagonals of [%d * %d] grid: %d", n, n, sum)
}
