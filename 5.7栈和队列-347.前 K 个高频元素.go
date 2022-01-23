## 栈和队列
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
即在数组中找到k个出现次数最多的元素

package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	// 方法1 用map存储每个元素的个数 然后用内置函数排序
	res := []int{}
	mapNum := map[int]int{}
	// 遍历数组 存储每个元素个数
	for _, v := range nums {
		mapNum[v]++
	}
	// 把上述map的key遍历出 存储到res
	for key, _ := range mapNum {
		res = append(res, key)
	}
	// 用内置函数 排序
	sort.Slice(res, func(i, j int) bool {
		// 使用每个元素的次数排序 大->小
		return mapNum[res[i]] > mapNum[res[j]]
	})
	return res[:k] // 返回前k个元素
}
func main() {
	fmt.Println(topKFrequent([]int{2, 2, 2, 1, 3, 3}, 2)) // [2 3]
}
