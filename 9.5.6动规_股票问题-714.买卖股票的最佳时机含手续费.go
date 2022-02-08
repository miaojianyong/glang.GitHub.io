## 动规_股票问题 买卖多次 每次有手续费

一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用

输入：prices = [1, 3, 2, 8, 4, 9], fee = 2
输出：8
解释：能够达到的最大利润:  
在此处买入 prices[0] = 1
在此处卖出 prices[3] = 8
在此处买入 prices[4] = 4
在此处卖出 prices[5] = 9
总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8
即每次交易后需要去掉相应的手续费

// 动规
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n) // dp 数组
	dp[0][0] = -prices[0]   // 初始化
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])     // 买入
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee) // 卖出
	}
	return dp[n-1][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 贪心
func maxProfit(prices []int, fee int) int {
    var minBuy int = prices[0] //第一天买入
    var res int
    for i:=0;i<len(prices);i++{
        //如果当前价格小于最低价，则在此处买入
        if prices[i]<minBuy{
            minBuy=prices[i]
        }
        //如果以当前价格卖出亏本，则不卖，继续找下一个可卖点
        if prices[i]>=minBuy&&prices[i]-fee-minBuy<=0{
            continue
        }
        //可以售卖了
        if prices[i]>minBuy+fee{
            //累加每天的收益
            res+=prices[i]-minBuy-fee
            //更新最小值（如果还在收获利润的区间里，表示并不是真正的卖出，而计算利润每次都要减去手续费，所以要让minBuy = prices[i] - fee;，这样在明天收获利润的时候，才不会多减一次手续费！）
            minBuy=prices[i]-fee
        }
    }
    return res
}
