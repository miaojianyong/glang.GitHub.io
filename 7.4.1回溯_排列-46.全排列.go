## 回溯_排列
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:
输入: [1,2,3]
输出: [ [1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1] ]
即输出数组的所有排列方式集合

func permute(nums []int) [][]int {
    res := [][]int{}
	path := []int{}
	used := make(map[int]bool)
	var backTrack func()
	backTrack = func() {
		// 递归终止条件
		if len(path) == len(nums) {
			/*tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)*/
			//return 可省略

			// 改成下述代码
			res = append(res, append([]int{}, path...))
		}
		// 遍历nums数组 从0开始
		for i := 0; i < len(nums); i++ {
			if used[nums[i]] { // nums中的元素使用过就跳过
				continue
			}
			used[nums[i]] = true // 每用一个元素 都标记一下
			path = append(path, nums[i])
			backTrack() // 递归
			// 回溯 包含标记的元素
			used[nums[i]] = false
			path = path[:len(path)-1]
		}
	}
	backTrack()
	return res
}
