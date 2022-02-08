## 动规_子序列 不连续

两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列
即子序列可以不连续 但是顺序不能颠倒

输入：text1 = "abcde", text2 = "ace" 
输出：3  
解释：最长公共子序列是 "ace" ，它的长度为 3 

func longestCommonSubsequence(text1 string, text2 string) int {
	t1 := len(text1)
	t2 := len(text2)
	dp := make([][]int, t1+1)
	for i := range dp {
		dp[i] = make([]int, t2+1)
	}
	for i := 1; i <= t1; i++ {
		for j := 1; j <= t2; j++ {
			// 如 test1上个元素 和 test2上个元素相等 表示找到一个相同的元素
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else { // 否则 让test1上个元素和test2当前元素的值 和 test1当前元素和test2下个元素的值 取最大值
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[t1][t2]
}

func max(x,y int)int{
    if x > y {
        return x
    }
    return y
}
