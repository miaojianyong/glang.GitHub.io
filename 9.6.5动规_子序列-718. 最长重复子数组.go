## 动规_子序列 连续

给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度

输入：
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出：3
解释：
长度最长的公共子数组是 [3, 2, 1]
即找到相关数组共同的 连续子序列的长度

// 动态规格 二维
func findLength(nums1 []int, nums2 []int) int {
	a, b := len(nums1), len(nums2)
	res := 0
	dp := make([][]int, a+1)
	for i := 0; i <= a; i++ {
		dp[i] = make([]int, b+1)
	}

	for i := 1; i <= a; i++ { // 遍历a
		for j := 1; j <= b; j++ { // 遍历b
			if nums1[i-1] == nums2[j-1] { // 如果两个数组的对应位置值相等
				dp[i][j] = dp[i-1][j-1] + 1 // 公式：记录
			}
			if dp[i][j] > res { // 找出最大的
				res = dp[i][j]
			}
		}
	}
	return res
}

// 动态规格 一维 执行时间更少 内存消耗也更低
func findLength2(nums1 []int, nums2 []int) int {
	dp := make([]int, len(nums2)+1)
	res := 0
	for i := 1; i <= len(nums1); i++ { // 数组1 前到后
		for j := len(nums2); j > 0; j-- { // 数组2 后到前
			if nums1[i-1] == nums2[j-1] { // 如两个数组对应值相等
				dp[j] = dp[j-1] + 1 // 子序长度+1
			} else { // 否则 子序长度从0开始
				dp[j] = 0
			}
			if dp[j] > res {
				res = dp[j]
			}
		}
	}
	return res
}
