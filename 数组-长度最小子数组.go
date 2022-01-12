## 数组
描述：在数组中找到 大于等于 目标值target的子数组
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。

// 暴力解法
func minSubArrayLen(target int, nums []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// 结果、子数组长度
	n, res := len(nums), math.MaxInt32
	// 外层循环从i开始 内层从i=j开始 即把数组从向到后挨个循环执行
	for i := 0; i < n; i++ {
		sum := 0 // 子数组和
		for j := i; j < n; j++ {
			sum += nums[j]     // 让当前元素逐个相加
			if sum >= target { // 子数组元素和 大于等于 目标值
				// j - i + 1 即表示当前符合的子数组长度
				res = min(res, j-i+1) // 把长度最小的给结果res
				break                 // 找到一组符合条件的子数组 就退出 进行下次循环
			}
		}
	}
	if res == math.MaxInt32 { // 如res等于默认值 表示没找到符合条件的子数组
		return 0
	}
	return res
}

// 滑动窗口
func minSubArrayLen2(target int, nums []int) int {
	// 起始位置、子数组和、结果
	startIndex, sum, res := 0, 0, math.MaxInt32 // 结果默认值也可用 len(nums)+1 即子数组长度 不可能大于数组长度
	for i := range nums {
		sum += nums[i]
		for sum >= target {
			subLength := i - startIndex + 1 // 当前子数组长度
			if subLength < res {            // 取最小的子数组长度
				res = subLength
			}
			// 当前子数组和 减去开始位置对应的值
			sum -= nums[startIndex]
			startIndex++ // 窗口向后移动
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}
