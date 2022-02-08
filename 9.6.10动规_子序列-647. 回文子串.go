## 动规_子序列 回文

一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目

输入：s = "abc"
输出：3
解释：三个回文子串: "a", "b", "c"
即返回字符串的回文个数

// 动规
func countSubstrings(s string) int {
    count := 0
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- { // 倒序 遍历字符串
		for j := i; j < n; j++ { // 正序 遍历i-n区间的字符串
			if s[i] == s[j] { // 如相等
				if j-i <= 1 { // 相等 或 相差1
					count++
					dp[i][j] = true
				} else if dp[i+1][j-1] { // 前开始 后一个到 后开始 前一个 区间内是否是回文
					count++
					dp[i][j] = true
				}
			}
		}
	}
	return count
}

// 双指针 内存消耗更少
func countSubstrings(s string) int {
    extend := func(s string, i, j, n int) int {
		res := 0
		// 两个指针 在字符串区间范围内 且相同情况下
		for i >= 0 && j < n && s[i] == s[j] {
			res++
			// 两个指针分别向左 向右移动
			i--
			j++
		}
		return res
	}
	count := 0
	n := len(s)
	for i := 0; i < n; i++ {
		count += extend(s, i, i, n)   // 以 i 为中心点
		count += extend(s, i, i+1, n) // 以 i和i+1 为中心点
	}
	return count
}
