package main

import (
	"fmt"
	"math/big"
)

func main() {
	digitLen := 1000
	num1 := big.NewInt(1)
	num2 := big.NewInt(1)
	index := 2

	for {
		add := new(big.Int)
		add = add.Add(num1, num2)
		num1, num2 = num2, add
		index++

		str := add.String()
		if len(str) == digitLen {
			fmt.Println("Index of first 1000 digit Fibonacci number:", index)
			break
		}
	}

}
