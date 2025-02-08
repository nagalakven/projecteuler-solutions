package main

import "fmt"

func main() {
	num1, num2, evenSum, add := 1, 2, 2, 0
	maxVal := 4000000

	for add < maxVal {
		add = num1 + num2
		num1, num2 = num2, add

		if add%2 == 0 {
			evenSum += add
		}
	}

	fmt.Println("Even-valued sum:", evenSum)
}
