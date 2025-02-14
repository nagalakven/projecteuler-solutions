package main

import (
	"fmt"
	"math"
	"strconv"
)

func findUpperBound(maxPower int) int {
	d := 1

	for {
		maxSum := d * maxPower
		minNumber := int(math.Pow(float64(10), float64(d-1)))

		if maxSum < minNumber {
			break
		}

		d++
	}

	return (d - 1) * maxPower
}

func main() {
	power := 5
	limit := 10 //minimum 2 digit number to get sum
	resultSum := 0
	fifthPowers := make([]int, limit)

	//pre-process 5th powers of digits
	for i := 1; i < limit; i++ {
		fifthPowers[i] = int(math.Pow(float64(i), float64(power)))
	}

	n := findUpperBound(fifthPowers[limit-1])

	for i := n; i > limit; i-- {
		numStr := strconv.Itoa(i)
		expoSum := 0
		for _, char := range numStr {
			val := int(char - '0')
			expoSum += fifthPowers[val]

			if expoSum > i {
				break
			}
		}

		if expoSum == i {
			resultSum += i
		}
	}

	fmt.Println("Digit fifth power sum:", resultSum)
}
