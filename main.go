package main

import "fmt"

func main() {
	var m, n int
	fmt.Scanln(&n, &m)
	var dp = make([]int, n+1)
	dp[0] = 1
	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			if j >= i {
				dp[j] += dp[j-i]
			}
		}
	}
	fmt.Print(dp[n])
}
