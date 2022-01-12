## 数组
描述：
  移除数组中指定的元素 返回操作后的数组长度

func removeElement(nums []int, val int) int {
	//length := len(nums)
	//res := 0 // 数组长度
	/*for i := 0; i < length; i++ { // 遍历数组
		if nums[i] != val { // 当前元素 不等于 目标值
			nums[res] = nums[i] // 把当前元素 给数组
			res++
		}
	}*/

	/*for _, v := range nums {
		if v != val {
			nums[res] = v
			res++
		}
	}*/
	//return res

	// 双指针法
	left, right := 0, len(nums)
	for left < right { // 左边小于右边
		if nums[left] == val { // 如当前值 等于 目标值
			nums[left] = nums[right-1] // 把最后一个值 给当前值
			right--                    // 向左移
		} else { // 否则 就是当前值 不等于 目标值
			left++ // 向右移 继续找
		}
	}
	//如 n = [1,2,3] val = 1 // 找到
	//n = [3,2] 即把最后的值 覆盖当前值 然后数组长度-1
	return left
}
