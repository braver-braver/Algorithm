package dp_test

import "testing"

// 输入：stones = [2,7,4,1,8,1]
// 输出：1
func lastStoneWeightII(stones []int) int {
	var sum int
	for _, v := range stones {
		sum += v
	}
	var target = sum / 2
	var dp = make([]int, target+1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			if dp[j] < dp[j-stones[i]]+stones[i] {
				dp[j] = dp[j-stones[i]] + stones[i]
			}
		}
	}
	return sum - 2*dp[target]
}

func TestLastStonrWeightII(t *testing.T) {
	t.Log(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
}
