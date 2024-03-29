package dp_test

import "testing"

// 输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
// 输出：4
func findMaxForm(strs []string, m int, n int) int {
	var dp = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// dp[i][j] = dp[i-count[0]][j-count[1]] + 1
	for i := range strs {
		s := strs[i]
		var zeroCount, oneCount int
		var sb = []byte(s)
		for i := 0; i < len(s); i++ {
			if sb[i] == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		}
		for i := m; i >= zeroCount; i-- {
			for j := n; j >= oneCount; j-- {
				if dp[i][j] < dp[i-zeroCount][j-oneCount] + 1 {
					dp[i][j] = dp[i-zeroCount][j-oneCount] + 1
				}
			}
		}
	}
	return dp[m][n]
}

func TestFindMaxForm(t *testing.T) {
	t.Log(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}
