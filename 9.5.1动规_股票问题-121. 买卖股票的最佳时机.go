## 动规_股票问题 -- 只能买卖1次

给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0

输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票

package main

import (
	"fmt"
	"math"
)

// 暴力枚举法 -- 执行超时
func maxProfit(prices []int) int {
	res := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			// res结果集 和 (上个值-当前值)的结果 取最大值
			res = getmax(res, prices[j]-prices[i])
		}
	}
	return res
}

// 贪心算法 -- 执行效率 高于动规
func maxProfit2(prices []int) int {
	low := math.MaxInt32
	res := 0
	for i := 0; i < len(prices); i++ {
		// 取出最小值
		low = getmin(low, prices[i])
		// 当前值-最小值 找出最大值
		res = getmax(res, prices[i]-low)
	}
	return res

	// 优化版 -- 效率最高
	//min := math.MaxInt
	//max := 0
	//for i := 0; i < len(prices); i++ {
	//	if prices[i] < min {
	//		min = prices[i]
	//	} else if prices[i] - min > max {
	//	max = prices[i] - min
	//	}
	//}
	//return max
}

func maxProfit3(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < length; i++ {
		dp[i][0] = getmax(dp[i-1][0], -prices[i])
		dp[i][1] = getmax(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[length-1][1]
}

func getmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getmin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))  // 5
	fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4})) // 5
	fmt.Println(maxProfit3([]int{7, 1, 5, 3, 6, 4})) // 5
}
