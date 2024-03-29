package dp_test

import "testing"

func uniquePaths(m int, n int) int {
	// return dfs(1, 1, m, n)
	var dp = make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func dfs(i, j int, m, n int) int {
	if i > m || j > n {
		return 0
	}
	if i == m && j == n {
		return 1
	}
	return dfs(i+1, j, m, n) + dfs(i, j+1, m, n)
}
func TestUniquePaths(t *testing.T) {
	t.Log(uniquePaths(3, 7))
}
