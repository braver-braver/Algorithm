package dp_test

import "fmt"

// 4 5
// 1 2
// 2 4
// 3 4
// 4 5
// 输出示例
// 10
func totalBag() {
	var m, n int
	fmt.Scanln(&m, &n)
	var weights = make([]int, m)
	var values = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scanln(&weights[i], &values[i])
	}
	var dp = make([]int, n+1)
	for i := 0; i < m; i++ {
		for j := weights[i]; j <= n; j++ {
			if dp[j] < dp[j-weights[i]]+values[i] {
				dp[j] = dp[j-weights[i]] + values[i]
			}
		}
	}
	fmt.Println(dp[n])
}
