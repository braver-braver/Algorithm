package main

import "fmt"

func main() {
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
