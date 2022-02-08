## 动规_完全背包

一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是

输入：n = 12
输出：3 
解释：12 = 4 + 4 + 4

输入：n = 12
输出：3 
解释：12 = 4 + 4 + 4
即：返回和是目标值n的完全平方数的个数

package main

import (
	"fmt"
	"math"
)

// 版本一,先遍历物品, 再遍历背包
func numSquares1(n int) int {
	//定义
	dp := make([]int, n+1)
	// 初始化
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	// 遍历物品
	for i := 1; i <= n; i++ {
		// 遍历背包
		for j := i * i; j <= n; j++ {
			dp[j] = getMin1(dp[j], dp[j-i*i]+1)
		}
	}

	return dp[n]
}

// 版本二,先遍历背包, 再遍历物品 -- 执行超时
func numSquares2(n int) int {
	//定义
	dp := make([]int, n+1)
	// 初始化
	dp[0] = 0
	// 遍历背包
	for j := 1; j <= n; j++ {
		//初始化
		dp[j] = math.MaxInt32
		// 遍历物品
		for i := 1; i <= n; i++ {
			if j >= i*i {
				dp[j] = getMin1(dp[j], dp[j-i*i]+1)
			}
		}
	}
	return dp[n]
}

func getMin1(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	fmt.Println(numSquares1(12)) // 3 4+4+4
}
