## 贪心_序列

如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。
第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也视作摆动序列。

输入：nums = [1,7,4,9,2,5]
输出：6
解释：整个序列均为摆动序列，各元素之间的差值为 (6, -3, 5, -7, 3) 
即数组中后面一个元素 减去前面一个元素的差是 正负树交替的 如上所示

func wiggleMaxLength(nums []int) int {
	var count int      // 记录峰值的个数
	var curDiff int    // 当前一对的差值
	var preDiff int    // 前一对的差值
	count = 1          // 默认序列最右边有一个峰值 即最后一个元素
	if len(nums) < 2 { // 题目数组长度最小是1
		return count
	}
	for i := 0; i < len(nums)-1; i++ { // 最右边的峰值 上述处理了 故-1操作
		// 后面的元素 减去 当前元素 即为当前差值
		curDiff = nums[i+1] - nums[i]
		// 如 (当前差值 大于0 且 前一个差值 小于等于0) 或 (当前差值 小于0 且 前一个差值 大于等于0)
		if (curDiff > 0 && preDiff <= 0) || (curDiff < 0 && preDiff >= 0) {
			preDiff = curDiff // 前一个差值 就等于 当前差值
			count++           // 峰值+1
		}
	}
	return count
}
