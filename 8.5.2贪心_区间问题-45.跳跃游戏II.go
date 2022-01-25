贪心_区间问题

给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
假设你总是可以到达数组的最后一个位置

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

func jump(nums []int) int {
    if len(nums) == 1 {
		return 0
	}
	// 当前覆盖的最远下标
	curIndex := 0
	// 下步覆盖的最远下标
	nextIndex := 0
	res := 0              // 步数
	for i := range nums { // i是nums中每个元素的索引
		// 更新下步覆盖的下标
		//nextIndex = func(a, b int) int {
		//	if a > b {
		//		return a
		//	}
		//	return b
		//}(i+nums[i], nextIndex)

		// 用内置函数找最大值
		nextIndex = int(math.Max(float64(i+nums[i]), float64(nextIndex)))

		// 如果数组当前下标 等于 当前覆盖的最远下标
		if i == curIndex {
			// 如 未覆盖到最后 步数+1
			if curIndex != len(nums)-1 {
				res++
				// 更新下标 到下步
				curIndex = nextIndex
				// 如 下步下标大于等于 数组最后的下标 退出
				if nextIndex >= len(nums)-1 {
					break
				}
			}
		}
	}
	return res
}
