package main

import (
	"fmt"
	"main/utils"
)

func main() {
	n := 2000000
	sum := 0

	for i := 2; i < n; i++ {
		if utils.IsPrime(i) {
			sum += i
		}
	}

	fmt.Println("Prime sum:", sum)
}
