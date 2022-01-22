## 哈希表
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]] 即(-1)+(-1)+2=0 (-1)+0+1=0
即输出数组中3个元素的和为0的组合

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// 哈希法
	//res := [][]int{}
	//count := map[int]int{}   // 存储每个元素和其个数
	//for _, v := range nums { // 遍历数组 把每个元素存储到map
	//	count[v]++ // key是元素值 value是该元素的个数
	//}
	//uniqNums := []int{}      // 存储去重后的数组
	//for key := range count { // 遍历map
	//	uniqNums = append(uniqNums, key) // 把map每个key存储到 去重数组
	//}
	//sort.Ints(uniqNums)                  // 排序 升序
	//for i := 0; i < len(uniqNums); i++ { // 遍历去重数组
	//	if uniqNums[0] > 0 { // 如数组的第1个元素大于0 无法组成三元组 跳过
	//		continue
	//	}
	//	// 符合三元组的形式 1.0+0+0 2.a+a+b/a+b+b 3.a+b+c
	//	// 1.0+0+0 即3个0*3等于0 且 该元素在map中的个数大于等于3个
	//	if uniqNums[i]*3 == 0 && count[uniqNums[i]] >= 3 {
	//		// 条件成立 把这3个元素 追加到结果集
	//		res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
	//	}
	//	for j := i + 1; j < len(uniqNums); j++ { // 遍历i后面的元素
	//		// 2.1 a+a+b 即2个a加上b等于0 且 a的个数在map中的个数大于1
	//		if uniqNums[i]*2+uniqNums[j] == 0 && count[uniqNums[i]] > 1 {
	//			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
	//		}
	//		// 2.2 a+b+b 即1个a加上2个b等于0 且 b的个数在map中的个数大于1
	//		if uniqNums[i]+uniqNums[j]*2 == 0 && count[uniqNums[j]] > 1 {
	//			res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
	//		}
	//		// 3. a+b+c 即3个不同的元素相加等于0
	//		c := -(uniqNums[i] + uniqNums[j]) // 即c等于 负的a+b
	//		// 如c大于后面1个数b[即一定大于前面的数a] 且 c在map中的个数大于0
	//		if c > uniqNums[j] && count[c] > 0 {
	//			res = append(res, []int{uniqNums[i], uniqNums[j], c})
	//		}
	//	}
	//}
	//return res

	// 双指针法
	// 思路：如[-3,-1,0,3,4] i是当前元素-3，left指向i的下个元素 即-1 right指向元素最后 即4
	// 1. 如a+b+c=0,那么把这3个元素 追加到结果集中
	// 2. 如a+b+c>0,即3个数的和大了 因是排序后的数组 故right指针应该向左移动 指向小点的元素
	// 3. 如a+b+c<0,即3个数的和小了 因是排序后的数组 故left指针应该向右移动 指向大点的元素
	sort.Ints(nums)
	res := [][]int{}
	n := len(nums)
	for i := 0; i < n-2; i++ { // 因为有了两个指针 故长度-2
		a := nums[i]
		if a > 0 { // 第1个元素大于0 无法组成三元组 故返回
			break
		}
		if i > 0 && a == nums[i-1] { // 发现相同的元素 就跳过 即去重
			continue
		}
		l, r := i+1, n-1 // 定义left和right指针的默认值
		//for l < r {      // 左指针小于右指针
		//	b, c := nums[l], nums[r] // 给n2 n3赋值
		//	// 1. 如符合 a+b+c=0 那么把这3个元素 追到到结果集
		//	if a+b+c == 0 {
		//		res = append(res, []int{a, b, c})
		//		// 左指针小于右指针 且 b的值等于l指针指向的元素
		//		for l < r && b == nums[l+1] { // 去重
		//			l++ // 让l指针 右移 遍历其他元素
		//		}
		//		// 左指针小于右指针 且 c的值等于r指针指向的元素
		//		for l < r && c == nums[r-1] { // 去重
		//			r-- // 让r指针 左移 遍历其他元素
		//		}
		// 		两个指针向中间靠
		//		l++
		//		r--
		//		// 2. 如 a+b+c>0 即3个数的和大了 故 r指针应该向左移动 指向小点的元素
		//	} else if a+b+c > 0 {
		//		r--
		//		// 3. 否则就是 a+b+c<0 即3个数的和小了 故 l指针应该向右移动 指向大点的元素
		//	} else {
		//		l++
		//	}
		//}
		for l < r {
			if nums[l]+nums[r]+nums[i] == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			} else if nums[l]+nums[r]+nums[i] > 0 {
				r--
			} else if nums[l]+nums[r]+nums[i] < 0 {
				l++
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums)) // [[-1 -1 2] [-1 0 1]]
}
