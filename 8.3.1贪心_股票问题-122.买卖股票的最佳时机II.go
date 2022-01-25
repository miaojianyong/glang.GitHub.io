## 贪心_股票问题

给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格

输入: prices = [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 
即根据上述数组买卖股票 使其利润最大

func maxProfit(prices []int) int {
    res := 0
	for i := 1; i < len(prices); i++ {
		res += max(0, prices[i]-prices[i-1])
	}
	return res
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

// 贪心 -- 优化
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
