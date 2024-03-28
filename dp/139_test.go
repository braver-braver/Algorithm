package dp_test

import "testing"

// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

// 示例 1：
// 输入: s = "leetcode", wordDict = ["leet", "code"]
// 输出: true
// 解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。

// 示例 2：
// 输入: s = "applepenapple", wordDict = ["apple", "pen"]
// 输出: true
// 解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
// 注意，你可以重复使用字典中的单词。

// 示例 3：
// 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
// 输出: false

type dict struct {
	val int
	m   map[string]struct{}
}

func wordBreak(s string, wordDict []string) bool {
	var ls = len(s)
	var dp = make([]dict, ls+1)
	for i := range dp {
		dp[i] = dict{
			val: -1,
			m:   make(map[string]struct{}),
		}
	}
	dp[0].val = 0
	for i := 0; i < len(wordDict); i++ {
		word := wordDict[i]
		for j := len(word); j <= ls; j++ {
			if dp[j-len(word)].val != -1 {
				if len(dp[j-len(word)].m) == 0 {
					dp[j].m[word] = struct{}{}
					dp[j].val = len(dp[j].m)
				} else {
					for k, v := range dp[j-len(word)].m {
						dp[j].m[k+word] = v
						dp[j].m[word+k] = v
						dp[j].val = len(dp[j].m)
					}
				}
			}
		}
	}
	if _, ok := dp[ls].m[s]; ok {
		return true
	}
	return false
}

func TestWordBreak(t *testing.T) {
	t.Log(wordBreak("applepenapple", []string{"apple", "pen"}))
}
