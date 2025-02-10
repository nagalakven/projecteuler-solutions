package main

import (
	"fmt"
	"main/utils"
)

func main() {
	n := 10000
	divSum := make(map[int]int)
	visited := make(map[int]bool) //mark visited pairs
	amicableNumSum := 0

	for i := 1; i < n; i++ {
		result := utils.DivisorsSum(i)
		if result != i {
			divSum[i] = result
		}
	}

	for key, value := range divSum {
		if foundVal, exists := divSum[value]; exists && foundVal == key && key != value {
			if !visited[key] && !visited[value] {
				amicableNumSum += key + value
				visited[key], visited[value] = true, true
			}
		}
	}

	fmt.Println("Amicable numbers sum:", amicableNumSum)
}
