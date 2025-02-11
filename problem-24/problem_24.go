package main

import (
	"fmt"
	"main/utils"
)

func main() {
	entry := 1000000
	targetIdx := entry - 1
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := len(digits)
	factorials := make([]int, n)
	numDigits := make([]int, n)

	for i := 0; i < n; i++ {
		factorials[i] = int((utils.Factorial(i)).Int64())
	}

	for i := 0; i < n; i++ {
		digitsLen := len(digits)
		factorial := factorials[digitsLen-1]

		idx := targetIdx / factorial
		numDigits[i] = digits[idx]

		targetIdx %= factorial
		digits = append(digits[:idx], digits[idx+1:]...)
	}

	result := 0

	for _, val := range numDigits {
		result = result*10 + val
	}

	fmt.Println("Millionth permutation is", result)
}
