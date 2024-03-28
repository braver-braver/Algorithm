package dp_test

import "fmt"

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

// 每次你可以爬至多m (1 <= m < n)个台阶。你有多少种不同的方法可以爬到楼顶呢？

// 注意：给定 n 是一个正整数。

// 输入描述
// 输入共一行，包含两个正整数，分别表示n, m
// 输出描述
// 输出一个整数，表示爬到楼顶的方法数。
// 输入示例
// 3 2
// 输出示例
// 3

func stairsII() {
	var m, n int
	fmt.Scanln(&m, &n)
	var dp = make([]int, n+1)
	dp[1] = 1
	for i := 0; i < m; i++ {
		for j := i + 1; j <= n; j++ {
			dp[j] += dp[j-i-1]
		}
	}
	fmt.Print(dp[n])
}
