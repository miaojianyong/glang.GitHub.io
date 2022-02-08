## 动规_完全背包

一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
假设每一种面额的硬币有无限个

输入：amount = 5, coins = [1, 2, 5]
输出：4
解释：有四种方式可以凑成总金额：
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
即：目标值是5，使用数组中的元素值，组合成和等于目标值的组合数

func change(amount int, coins []int) int {
    // 定义dp数组
	dp := make([]int, amount+1)
	// 初始化,0大小的背包, 当然是不装任何东西了, 就是1种方法
	dp[0] = 1
	// 遍历顺序
	for i := 0; i < len(coins); i++ { // 遍历物品
		for j := coins[i]; j <= amount; j++ { // 遍历背包
			// 推导公式
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}
