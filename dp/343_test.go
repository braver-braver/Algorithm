package dp_test

import (
	"testing"
)

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	var dp = make([]int, n+1)
	dp[2] = 1
	dp[1] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func TestIntegerBreak(t *testing.T) {
	t.Log(integerBreak(10))
}

