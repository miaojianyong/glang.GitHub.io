## 回溯_组合
一个 无重复元素 的整数数组 candidates 和一个目标整数 target，
找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。
可以按 任意顺序 返回这些组合

输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合
即根据数组的元素 找出能够组成和为目标值的组合 数组的每个元素可重复使用

package main

import (
	"fmt"
	"sort"
)

func combinationSum1(candidates []int, target int) [][]int {
	var trcak []int
	var res [][]int
	backtracking(0, 0, target, candidates, trcak, &res)
	return res
}
func backtracking(startIndex, sum, target int, candidates, trcak []int, res *[][]int) {
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
		trcak = append(trcak, candidates[i])                 // 更新路径
		sum += candidates[i]                                 // 更新sum
		backtracking(i, sum, target, candidates, trcak, res) // 递归 不用i+1因为可以重复读取当前数
		// 回溯
		trcak = trcak[:len(trcak)-1]
		sum -= candidates[i]
	}
}

// 优化 回溯
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates) // 排序数组
	n := len(candidates)
	group := []int{}
	var dfs func(int, int)

	dfs = func(idx int, sum int) {
		// 终止条件
		if target == sum {
			// tmp := make([]int,len(group))
			// copy(tmp,group)
			// res = append(res,tmp)
			// 简写
			res = append(res, append([]int{}, group...))
			return
		}
		for i := idx; i < n; i++ {
			if candidates[i]+sum > target {
				return
			} else if i < n-1 && candidates[i+1] == candidates[i] {
				continue
			}
			// 做选择
			group = append(group, candidates[i])
			sum += candidates[i]
			// 递归
			dfs(i, sum)
			// 删除选择
			group = group[:len(group)-1]
			sum -= candidates[i]
		}
	}
	dfs(0, 0)
	return res
}

func main() {
	fmt.Println(combinationSum1([]int{2, 3, 6, 7}, 7)) // [[2 2 3] [7]]
	fmt.Println(combinationSum2([]int{2, 3, 6, 7}, 7)) // [[2 2 3] [7]]
}
