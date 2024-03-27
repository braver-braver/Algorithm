package dp_test

import (
	"fmt"
	"testing"
)

// 输入示例
// 6 1
// 2 2 3 1 5 2
// 2 3 1 5 4 3
// 输出示例
// 5

func TestMaxValue(t *testing.T) {
	var kind, capacity int
	fmt.Scanf("%d %d", &kind, &capacity)
	var weight, value []int
	weight = make([]int, kind)
	value = make([]int, kind)
	for i := 0; i < kind; i++ {
		fmt.Scan(&weight[i])
	}
	for i := 0; i < kind; i++ {
		fmt.Scan(&value[i])
	}

	var dp = make([]int, capacity+1)
	dp[0] = 0
	for i := 0; i < kind; i++ {
		for j := capacity; j >= weight[i]; j-- {
			if dp[j] < dp[j-weight[i]]+value[i] {
				dp[j] = dp[j-weight[i]] + value[i]
			}
		}
	}
	print(dp[capacity])
}

func TestBagMaxValue2(t *testing.T) {
	var kind, capacity int
	fmt.Scanf("%d %d", &kind, &capacity)
	var weight, value []int
	weight = make([]int, kind)
	value = make([]int, kind)
	for i := 0; i < kind; i++ {
		fmt.Scan(&weight[i])
	}
	for i := 0; i < kind; i++ {
		fmt.Scan(&value[i])
	}
	var dp = make([][]int, kind)
	for i := 0; i < kind; i++ {
		dp[i] = make([]int, capacity+1)
	}
	//for j := capacity; j >= weight[0]; j-- {
	//	dp[0][j] = dp[0][j-weight[0]] + weight[0]
	//}
	for j := 1; j <= capacity; j++ {
		if j >= weight[0] {
			dp[0][j] = value[0]
		}
	}
	for i := 1; i < kind; i++ {
		for j := 1; j <= capacity; j++ {
			if j >= weight[i] {
				if dp[i-1][j] < dp[i-1][j-weight[i]]+value[i] {
					dp[i][j] = dp[i-1][j-weight[i]] + value[i]
				} else {
					dp[i][j] = dp[i-1][j]
				}
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	fmt.Println(dp[kind-1][capacity])
}
