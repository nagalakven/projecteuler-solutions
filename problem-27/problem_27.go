package main

import (
	"fmt"
	"main/utils"

	"github.com/fxtlabs/primes"
)

func consecutivePrimesCount(a, b int) int {
	primesCnt := 0

	for {
		result := (primesCnt * primesCnt) + (a * primesCnt) + b
		if !primes.IsPrime(result) {
			break
		}

		primesCnt++
	}

	return primesCnt
}

func main() {
	n := 1000
	bVals := utils.FindAllPrimes(n)
	cnt, maxCnt, maxA, maxB := 0, 0, 0, 0

	for a := -n + 1; a < n; a++ {
		for _, b := range bVals {
			cnt = consecutivePrimesCount(a, b)

			if cnt > maxCnt {
				maxCnt = cnt
				maxA, maxB = a, b
			}
		}
	}

	fmt.Println("Quadratic Primes - Coefficients product:", maxA*maxB)
}
