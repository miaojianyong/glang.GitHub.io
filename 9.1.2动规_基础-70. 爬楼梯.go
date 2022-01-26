## 动规_基础

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢

输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
即2阶楼梯 可1次爬1阶 可1次爬2阶 故有两种方法爬完

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1) // 1.定义dp数组
	// 2.初始化
	dp[1] = 1
	dp[2] = 2
	// 3. 遍历 从前向后
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 优化空间复杂度
func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	dp := [3]int{}
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		sum := dp[1] + dp[2]
		dp[1] = dp[2]
		dp[2] = sum
	}
	return dp[2]

	// 简写
	// 动规 简写
	//p, q, r := 0, 0, 1
	//for i := 0; i < n; i++ {
	//	p = q
	//	q = r
	//	r = p + q
	//}
	//return r
}
