## 动规_子序列 不连续

一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4
即：找到数组中最长的递增子序列的长度

// 动规
func lengthOfLIS2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max300(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 动规 + 二分查找 执行效率更高
func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		// 如当前dp是空的 或 当前dp中的最后一个元素 小于 当前元素
		if len(dp) == 0 || dp[len(dp)-1] < num {
			dp = append(dp, num) // 就把当前元素追加到dp中
		} else { // 二分查找
			l, r := 0, len(dp)-1
			tmp := r
			for l <= r {
				mid := (l + r) >> 1
				if dp[mid] >= num {
					tmp = mid
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			dp[tmp] = num
		}
	}
	return len(dp)
}
