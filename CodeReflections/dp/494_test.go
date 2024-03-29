package dp_test

import (
	"fmt"
	"testing"
)

// 示例 1：

// 输入：nums = [1,1,1,1,1], target = 3
// 输出：5
// 解释：一共有 5 种方法让最终目标和为 3 。
// -1 + 1 + 1 + 1 + 1 = 3
// +1 - 1 + 1 + 1 + 1 = 3
// +1 + 1 - 1 + 1 + 1 = 3
// +1 + 1 + 1 - 1 + 1 = 3
// +1 + 1 + 1 + 1 - 1 = 3
// 示例 2：

// 输入：nums = [1], target = 1
// 输出：1
func findTargetSumWays(nums []int, target int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if target < 0 {
		if -target > sum {
			return 0
		}
	} else {
		if target > sum {
			return 0
		}
	}
	target = sum + target
	if target%2 != 0 {
		return 0
	}
	target = target / 2
	var dp = make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	fmt.Printf("%v", dp)
	return dp[target]
}

func TestFindTargetSumWays(t *testing.T) {
	t.Log(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
