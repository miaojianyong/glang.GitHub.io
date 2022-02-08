## 动规_子序列 连续

一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度

输入：nums = [1,3,5,4,7]
输出：3
解释：最长连续递增序列是 [1,3,5], 长度为3。
尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开
即找到的子序列 必须是连续的，就是各元素是挨着的

// 动规
func findLengthOfLCIS(nums []int) int {
    if len(nums) == 0 {
		return 0
	}
	res := 1
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1 // 初始化dp数组 即连续子序最少为1
	}
	// 遍历nums数组 因要当前元素i和后面元素i+1比较 故len(nums)-1
	for i := 0; i < len(nums)-1; i++ {
		// 如 后面1个元素 大于当前元素
		if nums[i+1] > nums[i] {
			dp[i+1] = dp[i] + 1 // 记录连续子序
		}
		if dp[i+1] > res {
			res = dp[i+1] // 把dp的长度给结果res
		}
	}
	return res
}

// 动规 -- 优化执行时间
func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
	ret := 1

	for i := 1; i < n; i++ {
		dp[i] = 1
		if nums[i-1] < nums[i] {
			dp[i] = dp[i-1] + 1
		}
		ret = max(ret, dp[i])
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 贪心 -- 内存消耗相对少
func findLengthOfLCIS(nums []int) int {
    	if len(nums) == 0 {
		return 0
	}
	res := 1 // 连续子序最少为1
	count := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] { // 符合条件 count+1
			count++
		} else { // 否则 就不是连续的 count=1 重新开始
			count = 1
		}
		if count > res {
			res = count
		}
	}
	return res
}
