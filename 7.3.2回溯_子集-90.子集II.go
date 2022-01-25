## 回溯_子集
一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）
解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列

输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
即返回数组所有子集 且 不能包含重复的子集

func subsetsWithDup(nums []int) [][]int {
    res := [][]int{}
	sort.Ints(nums) // 排序
	// temp 存储一个子集
	var dfs func(temp, nums []int, start int)
	dfs = func(temp, nums []int, start int) {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		res = append(res, tmp)
		// 遍历nums
		for i := start; i < len(nums); i++ {
			// 去重
			if i > start && nums[i] == nums[i-1] { // 即区别上个子集问题 这里添加去重
				continue
			}
			temp = append(temp, nums[i])
			dfs(temp, nums, i+1)      // 递归
			temp = temp[:len(temp)-1] // 回溯
		}
	}
	dfs([]int{}, nums, 0)
	return res
}
