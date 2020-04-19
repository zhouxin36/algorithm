package main

import "fmt"

func main() {
	fmt.Println(numOfArrays(2, 3, 1) == 6)
	fmt.Println(numOfArrays(5, 2, 3) == 0)
	fmt.Println(numOfArrays(9, 1, 1) == 1)
	fmt.Println(numOfArrays(50, 100, 25) == 34549172)
	fmt.Println(numOfArrays(37, 17, 7) == 418930126)
}
func numOfArrays(n int, m int, k int) int {
	dp := [55][105][55]int{}
	dp[0][0][0] = 1
	mod := int(1e9 + 7)
	for i := 1; i <= m; i++ {
		dp[0][i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for a := 1; a <= k; a++ {
				dp[i][j][a] = (dp[i-1][j][a]-dp[i-1][j-1][a])*j%mod + dp[i-1][j-1][a-1]
				dp[i][j][a] = (dp[i][j][a]%mod + mod) % mod
				dp[i][j][a] += dp[i][j-1][a]
				dp[i][j][a] %= mod
			}
		}
	}
	return dp[n][m][k]
}
