## 动规_子序列 连续

一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6
即：找到数组中 连续子序的最大和

// 暴力解法
func maxSubArray(nums []int) int {
	res := math.MinInt64
	for i := 0; i < len(nums); i++ { // 设置起始位置
		sum := 0
		for j := i; j < len(nums); j++ { // 每次从起始位置i开始遍历 寻找最大值
			sum += nums[j]
			if sum > res {
				res = sum
			}
		}
	}
	return res
}

// 贪心解法
func maxSubArray2(nums []int) int {
	maxSum := nums[0]                // 第1个元素 给变量
	for i := 1; i < len(nums); i++ { // 从第2个元素开始遍历
		// 如果 当前元素 + 上个元素 大于 当前元素
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1] // 当前元素 = 当前元素+上个元素
		}
		// 如果当前元素 大于 最大和
		if nums[i] > maxSum {
			maxSum = nums[i] // 把当前元素 给 最大和
		}
	}
	return maxSum
}

// 动规
func maxSubArray(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    max := func(a,b int)int{
        if a > b {
            return a
        }
        return b
    }
    n := len(nums)
    dp := make([]int,n)
    dp[0] = nums[0]
    res := nums[0]
    for i:=1;i<n;i++{
        dp[i] = max(dp[i-1]+nums[i],nums[i])
        res = max(res,dp[i])
    }
    return res
}
