## 回溯_排列

一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列
示例 1：
输入：nums = [1,1,2]
输出： [[1,1,2], [1,2,1], [2,1,1]]
示例 2：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
即输出数组使用排列 且 不能重复

package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums) // 排序后方便去重操作
	res := [][]int{}
	path := []int{}
	used := make(map[int]bool)
	var backTrack func()
	backTrack = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
		}
		for i, v := range nums {
			if used[i] {
				continue
			}
			// 去重
			if i > 0 && !used[i-1] && v == nums[i-1] {
				continue
			}
			path = append(path, v)
			used[i] = true
			backTrack() // 递归
			// 回溯
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backTrack()
	return res
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))    // [[1 1 2] [1 2 1] [2 1 1]]
	fmt.Println(permuteUnique([]int{3, 3, 0, 3})) // [[0 3 3 3] [3 0 3 3] [3 3 0 3] [3 3 3 0]]
}
