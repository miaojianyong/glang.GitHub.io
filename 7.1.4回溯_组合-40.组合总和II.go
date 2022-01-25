## 回溯_组合
集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用 一次
注意：解集不能包含重复的组合

输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
即使用上述数组中的元素，组合和为目标值的组合，且数组中的每个元素只能用一次，组合不能重复

func combinationSum2(candidates []int, target int) [][]int {
	var trcak []int
	var res [][]int
	var histoy map[int]bool
	histoy = make(map[int]bool)
	sort.Ints(candidates) // 排序
	backtracking(0, 0, target, candidates, trcak, &res, histoy)
	return res
}
func backtracking(startIndex, sum, target int, candidates, trcak []int, res *[][]int, histoy map[int]bool) {
	if sum == target { // 终止条件
		tmp := make([]int, len(trcak))
		copy(tmp, trcak)
		*res = append(*res, tmp)
		return
	}
	if sum > target { // 终止条件
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] && histoy[i-1] == false {
			continue // 跳过使用过的元素
		}
		trcak = append(trcak, candidates[i]) // 更新路径
		sum += candidates[i]                 // 更新sum
		histoy[i] = true
		backtracking(i+1, sum, target, candidates, trcak, res, histoy) // 递归 不用i+1因为可以重复读取当前数
		// 回溯
		trcak = trcak[:len(trcak)-1]
		sum -= candidates[i]
		histoy[i] = false
	}
}

func main() {
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5)) // [[1 2 2] [5]]
}
