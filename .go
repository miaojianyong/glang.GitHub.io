## 数组

描述：
  找到 target 在数组中是索引 并返回
  没有找到 则返回-1

## 二分查找法
func search(nums []int, target int) int {
	left := 0              // 最左元素
	right := len(nums) - 1 // 最右元素
	for left <= right {
		//mid := (left + right) / 2 // 中间位置
		mid := (left + right) >> 1
		if nums[mid] == target { // 找到了
			return mid
		} else if nums[mid] > target { // 中间值大于目标值 去左侧找
			right = right - 1 // 向左移
		} else { // 否则 就是中间值 小于目标值 去右侧找
			left = left + 1 // 向右移
		}
	}
	return -1 // 没找到
}
