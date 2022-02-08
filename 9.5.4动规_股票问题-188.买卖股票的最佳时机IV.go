## 动规_股票问题 最多买卖 k 次

一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易

输入：k = 2, prices = [2,4,1]
输出：2
解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 || k == 0 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i:=0;i<len(prices);i++{
		dp[i] = make([]int, 2*k+1)
	}
	for i:=1;i<2*k+1;i+=2{ // 从1开始 奇数代表买入 初始都是-prices[0]
		dp[0][i] = -prices[0]
	}
	for i:=1;i<len(prices);i++{
        dp[i][0] = dp[i-1][0] // 无操作
		for j:=1;j<2*k+1;j++{
			if j % 2 == 1{ // 奇数 买入
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			}else{ // 偶数 卖出
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
			}
		}
	}
	return dp[len(prices)-1][2*k]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
