## 动规_01背包

一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等

输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 
即看数组是否可分割成两个子集 且这两个子集的和相对 即上述的 1+5+5 = 11

// 动规 二维
func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 { // 如果和为奇数 就不能分隔
		return false
	}
	sum /= 2 // 子集和
	// dp数组
	var dp [][]bool
	// 初始化
	dp = make([][]bool, len(nums)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= sum; j++ { // j是固定总量
			if j >= nums[i-1] { // 如果容量够用 则放入背包
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else { // 否则 不够用 维持前一个状态
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][sum]
}

// 动规 一维 -- 执行效率更高 消耗内存更少
func canPartition2(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, v := range nums {
		for j := target; j >= v; j-- {
			if dp[j] < dp[j-v]+v {
				dp[j] = dp[j-v] + v
			}
		}
	}
	return dp[target] == target
} 
