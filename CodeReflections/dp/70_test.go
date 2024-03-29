package dp_test

import "testing"

func climbStairs(n int) int {
	var dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func TestClimbStairs(t *testing.T) {
	t.Log(climbStairs(3))
}
