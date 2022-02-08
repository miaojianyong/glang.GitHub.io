## 动规_股票问题 买卖多次 卖出后 有冷冻期

一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​
设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)

输入: prices = [1,2,3,0,2]
输出: 3 
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	// 公式：
	// f[i][0]: 已买股票
	// f[i][1]: 没有有股票 且 在冷冻期
	// f[i][2]: 没有有股票 且 不在冷冻期
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[n-1][1], dp[n-1][2])
}
// 最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
