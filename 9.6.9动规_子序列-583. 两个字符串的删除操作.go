## 动规_子序列 编辑距离

两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。
每步 可以删除任意一个字符串中的一个字符

输入: word1 = "sea", word2 = "eat"
输出: 2
解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
即：删除两个字符串对应的字符，让其两个字符串相等 所用的最少步数

func minDistance(word1 string, word2 string) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	a, b := len(word1), len(word2)
	dp := make([][]int, a+1)
	for i := range dp {
		dp[i] = make([]int, b+1)
	}
	// 初始化
	for i := 0; i <= a; i++ { // w1为空时，w2全删
		dp[i][0] = i
	}
	for j := 0; j <= b; j++ { // w2为空时，w1全删
		dp[0][j] = j
	}
	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			if word1[i-1] == word2[j-1] { // 相对时
				dp[i][j] = dp[i-1][j-1]
			} else { // 否则 w1左移、w2左移的最小值 和 w1、w2都左移 去最小值
				// dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+2)
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1) // dp[i-1][j-1]+2一定比前面两个大 故可舍去
			}
		}
	}
	return dp[a][b]
}
