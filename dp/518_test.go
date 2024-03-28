package dp_test

import "testing"

// 输入：amount = 5, coins = [1, 2, 5]
// 输出：4

func change(amount int, coins []int) int {
	var dp = make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

func TestChangeII(t *testing.T) {
	t.Log(change(5, []int{1, 2, 5}))
}
