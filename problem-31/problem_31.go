package main

import "fmt"

func main() {
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	target := 200
	dp := make([]int, target+1)

	//initialize, only one way to make 0 (no coins)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= target; i++ {
			dp[i] += dp[i-coin]
		}
	}

	fmt.Println("Total ways:", dp[target])
}
