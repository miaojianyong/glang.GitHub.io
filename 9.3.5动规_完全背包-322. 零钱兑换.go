## 动规_完全背包

一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的

输入：coins = [1, 2, 5], amount = 11
输出：3 
解释：11 = 5 + 5 + 1
即：根据目标值，使用数组中的元素值，找到和为目标值的组合的元素个数

package main

import (
	"fmt"
	"math"
)

// 子问题拆分 -- 执行超时
func coinChange(coins []int, amount int) int {
	var def func(int) int
	def = func(amount int) int {
		// 目标金额为0是 所需的硬币数为0
		if amount == 0 {
			return 0
		}
		// 目标金额小于0时 无解返回-1
		if amount < 0 {
			return -1
		}
		// 求最小值 即最少硬币数
		res := math.MaxInt32 // 大点的值
		// 遍历硬币coins集合 取出每个硬币coin
		for _, coin := range coins {
			// 子问题 即当前目标金额 - 当前硬币
			subproblem := def(amount - coin)
			if subproblem == -1 { // 子问题无解 跳过
				continue
			}
			//在子问题中选择最优解 然后加一
			res = getMin(res, 1+subproblem)
		}
		if res == math.MaxInt32 {
			return -1
		} else {
			return res
		}
	}
	return def(amount)
}

// 备忘录 递归法
func coinChange2(coins []int, amount int) int {
	memo := make([]int, amount+1) // 这里每个都是0 可能需要更小的数要遍历memo再赋值
	var def func(int) int
	def = func(amount int) int {
		// 目标金额为0是 所需的硬币数为0
		if amount == 0 {
			return 0
		}
		// 目标金额小于0时 无解返回-1
		if amount < 0 {
			return -1
		}
		//查备忘录 防止重复计算
		if memo[amount] != 0 {
			return memo[amount]
		}
		// 求最小值 即最少硬币数
		res := math.MaxInt32 // 大点的值
		// 遍历硬币coins集合 取出每个硬币coin
		for _, coin := range coins {
			// 子问题 即当前目标金额 - 当前硬币
			subproblem := def(amount - coin)
			if subproblem == -1 { // 子问题无解 跳过
				continue
			}
			//在子问题中选择最优解 然后加一
			res = getMin(res, 1+subproblem)
		}
		// 把计算结果存入备忘录 不然无法跳出子问题的递归调用
		if res == math.MaxInt32 {
			memo[amount] = -1
		} else {
			memo[amount] = res
		}
		return memo[amount]
	}
	return def(amount)
}

// dp数组的定义：当目标金额为i时，至少需要dp[i]枚硬币凑出 -- 效率最高
func coinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	//初始化 数组大小为amount+1， 初始值也为amount+1
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	//base case
	dp[0] = 0
	//外层for循环遍历所有状态的所有取值
	for i := 1; i < len(dp); i++ {
		//内层循环求所有选择的最小值
		for _, coin := range coins {
			//子问题无解 跳过
			if i-coin < 0 {
				continue
			}
			dp[i] = getMin(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))  // 3
	fmt.Println(coinChange2(coins, amount)) // 3
	fmt.Println(coinChange3(coins, amount)) // 3
}
