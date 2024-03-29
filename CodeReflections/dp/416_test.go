package dp_test

import "testing"

// 输入：nums = [1,5,11,5]
// 输出：true
// 解释：数组可以分割成 [1, 5, 5] 和 [11] 。

func canPartition(nums []int) bool {
	var sum int
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2
	var dp = make([]int, sum+1)
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := sum; j >= nums[i]; j-- {
			if dp[j] < dp[j-nums[i]]+nums[i] {
				dp[j] = dp[j-nums[i]] + nums[i]
			}
		}
	}
	return dp[sum] == sum
}

func TestCanPartition(t *testing.T) {
	t.Log(canPartition([]int{1, 5, 11, 5}))
}
