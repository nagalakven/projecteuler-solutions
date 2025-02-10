package main

import (
	"fmt"
	"main/utils"
)

func main() {
	n := 28124
	abundant := []int{}
	canBe := make([]bool, n)

	//find sum of proper divisors
	for i := 1; i < n; i++ {
		divSum := utils.DivisorsSum(i)
		if divSum > i {
			abundant = append(abundant, i)
		}
	}

	abLen := len(abundant)
	//process list
	for i := 0; i < abLen; i++ {
		for j := i; j < abLen; j++ {
			sum := abundant[i] + abundant[j]
			if sum < n {
				canBe[sum] = true
			} else {
				//j increases going fwd
				break
			}
		}
	}

	//calculate sum of existing non-abundant numbers
	result := 0
	for i := 0; i < n; i++ {
		if !canBe[i] {
			result += i
		}
	}

	fmt.Println("Non-abundant sum:", result)
}
