## 贪心_区间问题

给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标
即根据数组的元素值 跳跃对应的步长 判断是否等跳到数组最后

func canJump(nums []int) bool {
    cover := 0 // 表示覆盖的下标
	if len(nums) == 1 {
		return true
	}
	// 从0下标开始 即第一个元素
	i := 0
	for i <= cover {
		// cover每次应在 原cover和当前的i+nums[1] 中取最大值
		cover = func(a, b int) int {
			if a > b {
				return a
			} else {
				return b
			}
		}(cover, i+nums[i])
		// 如果cover 大于等于 最后一个元素在的索引位置 返回true
		if cover >= len(nums)-1 {
			return true
		}
		// 更新下标
		i++
	}
	// 否则 就是走完了cover也不能 覆盖到最后元素 返回false
	return false
}
