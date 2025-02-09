package main

import "fmt"

func main() {
	n := 20
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= n; i++ {
		dp[0][i], dp[i][0] = 1, 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	fmt.Println("Lattic paths:", dp[n][n])
}
