package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func main() {

	n, result := 20, 1

	for i := 2; i <= n; i++ {
		result = lcm(result, i)
	}

	fmt.Println("Smallest multiple:", result)
}
