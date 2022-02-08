## 动规_股票问题 可以买卖多次

给定一个数组 prices ，其中 prices[i] 表示股票第 i 天的价格。
在每一天，你可能会决定购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以购买它，然后在 同一天 出售。
返回 你能获得的 最大 利润 

输入: prices = [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3
即：根据数组中的元素，买卖股票，达到最大利润

// 贪心 -- 内存消耗更少
func maxProfit(prices []int) int {
	var sum int
	// 从第2个元素开始遍历 因为2天是一个交易单位（当前元素 要减去 上个元素）
	for i := 1; i < len(prices); i++ {
		// 如 当前元素 减 上个元素 大于0
		if prices[i]-prices[i-1] > 0 {
			// 累加利润
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	status := make([]int, len(prices)*2)
	for i := range dp {
		dp[i] = status[:2]
		status = status[2:]
	}
	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[len(prices)-1][1]
}

// 优化内存消耗
func maxProfit(prices []int) int {
	//创建数组
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

func max(a,b int)int{
  if a > b {
    return a
  }
  return b
}
