## 动规_打家劫舍

你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。
这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的
即：数组中的元素是一个环，即头元素和尾元素相连，要获取单独的元素

思路：
情况一：考虑不包含首尾元素 
情况一：考虑不包含首尾元素
情况三：考虑包含尾元素，不包含首元素
注意我这里用的是"考虑"，例如情况三，虽然是考虑包含尾元素，但不一定要选尾部元素！ 对于情况三，取nums[1] 和 nums[3]就是最大的。
而情况二 和 情况三 都包含了情况一了，所以只考虑情况二和情况三就可以了

func rob(nums []int) int {
    if len(nums) == 1 { // 数组长度为1，返回第1个元素
        return nums[0]
    }
    if len(nums) == 2 { // 数组长度为2，返回数组最大的元素
        return max(nums[0], nums[1])
    }
    
    result1 := robRange(nums, 0)
    result2 := robRange(nums, 1)
    return max(result1, result2)
}

// 偷盗指定的范围
func robRange(nums []int, start int) int {
    dp := make([]int, len(nums))
    dp[1] = nums[start]
    
    for i := 2; i < len(nums); i++ {
        dp[i] = max(dp[i - 2] + nums[i - 1 + start], dp[i - 1])
    }
    
    return dp[len(nums) - 1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
