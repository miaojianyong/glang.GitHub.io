## 动规_子序列 回文

一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度
子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列

输入：s = "bbbab"
输出：4
解释：一个可能的最长回文子序列为 "bbbb"

// 动规 二维
func longestPalindromeSubseq(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1 // 初始化
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

// 动规 一维 执行时间更高 内存消耗更少
func longestPalindromeSubseq2(s string) int {
	n := len(s)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1 // 初始化
	}

	for i := n - 2; i >= 0; i-- {
		res := 0
		// 初始化1了 故从+1开始
		for j := i + 1; j < n; j++ {
			tmp := dp[j]
			if s[i] == s[j] {
				dp[j] = res + 2
			} else {
				if dp[j] < dp[j-1] {
					dp[j] = dp[j-1]
				}
			}
			res = tmp
		}
	}
	return dp[n-1]
}
