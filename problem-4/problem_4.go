package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(str string) bool {
	sLen := len(str)

	for i := 0; i < sLen/2; i++ {
		if str[i] != str[sLen-i-1] {
			return false
		}
	}

	return true
}

func main() {
	num1Max := 999
	num1Min, num2Min := 100, 100

	maxProd := 1

	for i := num1Max; i >= num1Min; i-- {
		for j := i; j >= num2Min; j-- {
			prod := i * j

			if prod <= maxProd {
				break
			}

			str := strconv.Itoa(prod)
			if isPalindrome(str) {
				maxProd = prod
			}
		}
	}

	fmt.Printf("Largest palindrom product: %d\n", maxProd)
}
