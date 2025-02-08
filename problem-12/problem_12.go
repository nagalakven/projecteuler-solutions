package main

import (
	"fmt"
	"main/utils"
	"math"
)

func main() {
	n := math.MaxInt
	triSum := 0
	maxDivisors := 500
	factors := make(map[int][]int)

	for i := 1; i <= n; i++ {
		triSum += i

		if utils.FindFactors(triSum, factors) > maxDivisors {
			fmt.Println("Highly divisible triangular number:", triSum)
			break
		}
	}
}
