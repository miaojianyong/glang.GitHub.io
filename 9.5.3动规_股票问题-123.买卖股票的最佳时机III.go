## 动规_股票问题 最多买卖两次

一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易

输入：prices = [3,3,5,0,0,3,1,4]
输出：6
解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 

func maxProfit(prices []int) int {
    // dp数组
	dp := make([][]int, len(prices))
	for i := range prices { // 存储5种状态
		dp[i] = make([]int, 5)
	}
	// 初始化 5种状态
	dp[0][0] = 0          // 无操作
	dp[0][1] = -prices[0] // 第1次 买入
	dp[0][2] = 0          // 第1次 卖出
	dp[0][3] = -prices[0] // 第2次 买入
	dp[0][4] = 0          // 第2次 卖出
	for i := 1; i < len(prices); i++ {
        // 对应5个状态 买卖时 下述两种方式
        // 1 买入时：当前状态的值 和 上个状态的值-当前股价 取最大值
        // 2 卖出时：当前状态的值 和 上个状态的值+当前股价 取最大值
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[len(prices)-1][4]
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
