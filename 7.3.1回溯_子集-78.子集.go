## 回溯_子集
一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
即输出上述数组的所有子集集合

func subsets(nums []int) [][]int {
	//res := make([][]int, 0)
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
			temp = append(temp, nums[i])
			dfs(temp, nums, i+1)      // 递归
			temp = temp[:len(temp)-1] // 回溯
		}
	}
	dfs([]int{}, nums, 0)
	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3})) // [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
}
