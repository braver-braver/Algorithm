package dp_test

import (
	"testing"
)

// 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。

// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。
// 例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

// 示例 1：

// 输入：n = 12
// 输出：3
// 解释：12 = 4 + 4 + 4
// 示例 2：

// 输入：n = 13
// 输出：2
// 解释：13 = 4 + 9

func numSquares(n int) int {
	var avaliable = make([]int, n)
	if n == 1 {
		return 1
	}
	var i int
	for i = 1; i <= n; i++ {
		v := i * i
		if v > n {
			break
		}
		avaliable[i-1] = v
	}
	avaliable = avaliable[:i-1]
	var dp = make([]int, n+1)
	for i := range dp {
		dp[i] = n + 1
	}
	dp[0] = 0
	for i := 0; i < len(avaliable); i++ {
		for j := avaliable[i]; j <= n; j++ {
			if dp[j-avaliable[i]]+1 < dp[j] {
				dp[j] = dp[j-avaliable[i]] + 1
			}
		}
	}
	return dp[n]
}

func TestNumSquares(t *testing.T) {
	t.Log(numSquares(4))
}
func Test_numSquares(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"一", args{
			12,
		}, 3},
		{
			"二",
			args{
				13,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSquares(tt.args.n); got != tt.want {
				t.Errorf("numSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
