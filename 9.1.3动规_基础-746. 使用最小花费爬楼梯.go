## 动规_基础

给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。
你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
请你计算并返回达到楼梯顶部的最低花费

输入：cost = [1,100,1,1,1,100,1,1,100,1]
输出：6
解释：你将从下标为 0 的台阶开始。
- 支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
- 支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
- 支付 1 ，向上爬一个台阶，到达楼梯顶部。
总花费为 6 。
即从数组中找到花费小的[小元素],去爬楼梯，返回爬完的最小花费

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)) // 定义dp数组
	// 初始化dp数组
	dp[0], dp[1] = cost[0], cost[1]
	// 遍历cost 前到后
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 优化空间复杂度
func minCostClimbingStairs2(cost []int) int {
	dp0 := cost[0]
	dp1 := cost[1]
	for i := 2; i < len(cost); i++ {
		temp := min(dp0, dp1) + cost[i]
		dp0 = dp1
		dp1 = temp
	}
	return min(dp0, dp1)
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20})) // 15
}
