## 动规_基础

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径

start   0  0  0  0   0   0
  0     0  0  0  0   0   0
  0     0  0  0  0   0   finish
输入：m = 3, n = 7
输出：28
即在 3*7的网格中 从start到finish有28种不同路径

// 抽象成 二叉树 解法 -- 运行超时
func uniquePaths(m int, n int) int {
	// 遍历二叉树 从1,1点开始 到m,n点
	var dfs func(i, j, m, n int) int
	dfs = func(i, j, m, n int) int {
		if i > m || j > n { // 越界
			return 0
		}
		if i == m && j == n { // 到了叶子节点 路径数+1
			return 1
		}
		// 向下 向右递归
		return dfs(i+1, j, m, n) + dfs(i, j+1, m, n)
	}
	return dfs(1, 1, m, n)
}

package main

import (
	"fmt"
	"math/big"
)

// 抽象成 二叉树 解法 -- 运行超时
func uniquePaths(m int, n int) int {
	// 遍历二叉树 从1,1点开始 到m,n点
	var dfs func(i, j, m, n int) int
	dfs = func(i, j, m, n int) int {
		if i > m || j > n { // 越界
			return 0
		}
		if i == m && j == n { // 到了叶子节点 路径数+1
			return 1
		}
		// 向下 向右递归
		return dfs(i+1, j, m, n) + dfs(i, j+1, m, n)
	}
	return dfs(1, 1, m, n)
}

// 动态规划 解法
func uniquePaths2(m int, n int) int {
	// dp数组
	dp := make([][]int, m)
	// 初始化dp数组
	for i := range dp { // 初始化每行 往下走
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ { // 初始化每列 往右走
		dp[0][j] = 1
	}
	// 遍历数组
	for i := 1; i < m; i++ { // 行
		for j := 1; j < n; j++ { // 列
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1] // 终点
}

// 数论 解法
func uniquePaths3(m int, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}

func main() {
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths2(3, 7))
	fmt.Println(uniquePaths3(3, 7))
}
