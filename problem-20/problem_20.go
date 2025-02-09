package main

import (
	"fmt"
	"main/utils"
)

func main() {
	n := 100

	result := utils.Factorial(n)
	resultStr := result.String()
	sum := 0

	for _, char := range resultStr {
		sum += int(char - '0')
	}

	fmt.Printf("Sum of digits of %d! : %d\n", n, sum)
}
