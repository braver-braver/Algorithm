package dp_test

import (
	"fmt"
	"testing"
)

// 输入：nums = [1,2,3], target = 4
// 输出：7

func combinationSum4(nums []int, target int) int {
	var dp = make([]int, target+1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i >= nums[j] {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	fmt.Printf("%v", dp)
	return dp[target]
}

func TestCombinationSum4(t *testing.T) {
	t.Log(combinationSum4([]int{1, 2, 3}, 4))
}
