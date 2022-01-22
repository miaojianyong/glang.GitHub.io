## 哈希表
描述：两个数组，编写一个函数来计算它们的交集。
示例：nums1 = [1,2,2,1] nums2 = [2,2,3]
输出：[2] 即两个数组的交集是2，为两个数组都存在的元素

package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	//numMap := make(map[int]int)
	//// 从nums1中取元素 作为key存储到map
	//// 相同的数仅即记录一次
	//// 如：[1,2,2] 记作：map[1:1,2:1]
	//for _, v := range nums1 {
	//	numMap[v] = 1
	//}
	//var res []int
	//// 从nums2中取元素
	//// 利用count>0，实现重复值只拿一次放入返回结果中
	//for _, v := range nums2 {
	//	// 如 取出的元素 在上述map中 且个数大于0 就添加到结果集
	//	// v是key,count是key对应的值 即2:1,2是v,1是count
	//	if count, ok := numMap[v]; ok && count > 0 {
	//		res = append(res, v)
	//		numMap[v]-- // 然后对应的个数减1
	//	}
	//}
	//return res

	//简写版
	table := make(map[int]int)
	res := []int{}
	for _, v := range nums1 {
		// 把当前元素 在上述map中的值设置为1，即默认值是0
		/*
			如：
			n := map[int]int{1: 2, 2: 3}
			num := n[3]
			fmt.Println(num, n[3]) // 0 0
		*/
		if val := table[v]; val < 1 { // val < 1 相当于去重 即设置过的就跳过了
			table[v] = 1
		}
	}
	for _, v := range nums2 {
		if val := table[v]; val == 1 { // 仅取val等于1的 即对应上述的设置
			res = append(res, v)
			table[v]++ // 添加结果集后 做+1操作 防止重复
		}
	}
	return res

	// 优化版，利用set，减少count统计
	//set := make(map[int]struct{}, 0)
	//res := make([]int, 0)
	//for _, v := range nums1 {
	//	if _, ok := set[v]; !ok { // 如果元素不在set中 就记录
	//		set[v] = struct{}{}
	//	}
	//}
	//for _, v := range nums2 {
	//	//如果存在于上一个数组中，则加入结果集，并清空该元素在set中的键值对
	//	if _, ok := set[v]; ok {
	//		res = append(res, v)
	//		delete(set, v) // 即防止重复
	//	}
	//}
	//return res
}

func main() {
	n1 := []int{1, 2, 2}
	n2 := []int{1, 1}
	fmt.Println(intersection(n1, n2)) // [1]

	n := map[int]int{1: 2, 2: 3}
	num := n[3]
	fmt.Println(num, n[3]) // 0 0
}
