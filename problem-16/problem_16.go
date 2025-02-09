package main

import (
	"fmt"
	"math/big"
)

func main() {
	n, pow := big.NewInt(2), 1000
	val := new(big.Int).Exp(n, big.NewInt(int64(pow)), nil)

	valStr := val.String()
	sum := 0
	for _, char := range valStr {
		sum += int(char - '0')
	}

	fmt.Println("Power digit sum:", sum)
}
