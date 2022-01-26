## 动规_基础

一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
即把一个整数拆分成至少2个数的和，找出这些数的最大乘积

// 动规
func integerBreak(n int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 1. 定义dp数组
	dp := make([]int, n+1)
	// 3. 初始化 2开始 即2=1+1 1*1=1
	dp[2] = 1
	// 4. 遍历 从3开始 即上述初始化了2
	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ { // j的范围 (1<=j<i)
			// 2. 公式
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

// 贪心 -- 内存消耗更少
func integerBreak2(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if n == 4 {
		return 4
	}
	res := 1
	for n > 4 {
		res *= 3
		n -= 3
	}
	res *= n
	return res
}
