## 哈希表
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1]

func twoSum(nums []int, target int) []int {
	// 1. 暴力法
	// for i := 0; i < len(nums); i++ { // 遍历数组 从第1个元素开始
	// 	for j := i + 1; j < len(nums); j++ { // 遍历数组 从第2个元素开始
	// 		// 如果数组中 有两个元素相加的和 等于 目标值 就返回其索引
	// 		if nums[i]+nums[j] == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }
	// return nil

    // 2. 使用map
	numMap := make(map[int]int, len(nums))
	for i, v := range nums {
		if preIndex, ok := numMap[target-v]; ok { // 如在map中能取到 目标值-当前值 的key对应的元素
			return []int{preIndex, i} // 就输出 上述差值对应的key，和当前值的索引
		} else { // 否则把当前值存储到map
			numMap[v] = i // 当前值作为key,索引作为value
		}
	}
	return []int{}
}
