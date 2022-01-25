## 贪心_股票问题

给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用

输入：prices = [1, 3, 2, 8, 4, 9], fee = 2
输出：8
解释：能够达到的最大利润:  
在此处买入 prices[0] = 1
在此处卖出 prices[3] = 8
在此处买入 prices[4] = 4
在此处卖出 prices[5] = 9
总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8
即每次买卖股票后需再-去手续费

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
