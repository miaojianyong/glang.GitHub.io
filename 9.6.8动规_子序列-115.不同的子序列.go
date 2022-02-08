## 动规_子序列 编辑距离

一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数

输入：s = "rabbbit", t = "rabbit"
输出：3
解释：
如下图所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
rabbbit
---- --

rabbbit
-- ----

rabbbit
--- ---
即如上所示，找出字符串的不同位置就算1次子序列

func numDistinct(s string, t string) int {
	sLen, tLen := len(s), len(t)
	// dp[i][j]：前 i 个字符的 s 子串中，出现前 j 个字符的 t 子串的次数
	dp := make([][]int, sLen+1) // 二维dp数组
	for i := range dp {
		dp[i] = make([]int, tLen+1)
	}
	for i := range dp { // 初始化
		dp[i][0] = 1
	}
	for i := 1; i < sLen+1; i++ { // 遍历dp矩阵，填表
		for j := 1; j < tLen+1; j++ { // dp[0][j]使用默认值0 故可不用初始化
			// 递推公式
			if s[i-1] == t[j-1] { // 相等 即在s中找到t
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else { // 否则 不相等
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[sLen][tLen] // 前 sLen 个字符的 s 串中，出现前 tLen 个字符的 t 串的次数
}

func numDistinct2(s string, t string) int {
	n1 := len(s)
	n2 := len(t)
	// 在s和t前面各加上空白字符
	s = " " + s
	t = " " + t

	f := make([][]int, n1+1)
	for i := range f {
		f[i] = make([]int, n2+1)
	}
	for i := 0; i <= n1; i++ {
		f[i][0] = 1
	}
	//for i := range f {
	//	f[i][0] = 1
	//}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if i < j {
				f[i][j] = 0
				continue
			}
			if s[i] == t[j] {
				f[i][j] = f[i-1][j] + f[i-1][j-1]
			} else {
				f[i][j] = f[i-1][j]
			}
		}
	}
	return f[n1][n2]
}

// 简洁版  执行时间更少 内存消耗更少
func numDistinct3(s string, t string) int {
	dp := make([]int,len(t)+1,len(t)+1)
	dp[0]=1
	for i:=1;i<=len(s);i++{
		for j:=len(t);j>=0;j--{
			if j>0 && s[i-1]==t[j-1] {dp[j] += dp[j-1]}
		}
	}
	return dp[len(t)]
} 
