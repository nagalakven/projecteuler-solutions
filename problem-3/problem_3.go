package main

import (
	"fmt"
	"math"

	"github.com/fxtlabs/primes"
)

func main() {
	num := 600851475143
	n := int(math.Sqrt(float64(num)))

	for i := n; i > 1; i-- {
		if num%i == 0 && primes.IsPrime(i) {
			fmt.Println("Largest prime factor:", i)
			break
		}
	}

	fmt.Println("No factor!!!")
}
