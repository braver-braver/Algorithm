package dp_test

import "testing"

func minCostClimbingStairs(cost []int) int {
	var dp = make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < len(dp); i++ {
		if dp[i-1]+cost[i-1] < dp[i-2]+cost[i-2] {
			dp[i] = dp[i-1] + cost[i-1]
		} else {
			dp[i] = dp[i-2] + cost[i-2]
		}
	}
	return dp[len(dp)-1]
}

func TestMinCostClimbingStairs(t *testing.T) {
	t.Log(minCostClimbingStairs([]int{10, 15, 20}))
	t.Log(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
