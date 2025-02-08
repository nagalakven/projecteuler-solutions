package main

import (
	"fmt"
	"main/utils"
)

func main() {
	cnt := 0
	num := 1

	for cnt < 10001 {
		num++
		if utils.IsPrime(num) {
			cnt++
		}
	}

	fmt.Println("10001st prime is:", num)
}
