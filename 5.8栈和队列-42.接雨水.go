## 栈和队列
输入：height=[1,0,2,1,3,1,0,1,2,0,1]
输出：7
即：每个元素的值表示柱子的高度，两个柱子直接可存雨水，如下述元素0-2之间可存雨水为1
-                   3
-         2    *        *   *   *   2
- 1   *        1        1   *   *       *   1
-     0                     0   1       0   
-----------------------------------------------

package main

import "fmt"

func trap(height []int) int {
	// 双指针法
	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	//var max func(a, b int) int
	//max = func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}

	//sum := 0
	//for i := 0; i < len(height); i++ {
	//	// 第1个柱子和最后1个柱子不接雨水
	//	if i == 0 || i == len(height)-1 {
	//		continue
	//	}
	//	lHig := height[i] // 记录左侧柱子的最高高度
	//	rHig := height[i] // 记录右侧柱子的最高高度
	//	// 右边柱子 找 当前i往右 的最大值
	//	for r := i + 1; r < len(height); r++ {
	//		if height[r] > rHig {
	//			rHig = height[r]
	//		}
	//	}
	//	// 左边柱子 找 当前[i-1]范围 的最大值
	//	for l := i - 1; l >= 0; l-- {
	//		if height[l] > lHig {
	//			lHig = height[l]
	//		}
	//	}
	//	// 取上述两个柱子的最小值 即低柱子的值 才是存雨水的值
	//	// 然后在减去 当前柱子的高度 就是存的雨水
	//	h := min(lHig, rHig) - height[i]
	//	if h > 0 { // 如上述h大于0 即存了雨水
	//		sum += h
	//	}
	//}
	//return sum

	//// 动态规格
	//// 当前雨水面积：[min(左边柱子的最高高度, 右边柱子的最高高度) - 当前柱子高度]*单位宽度[这里是1]
	//n := len(height)
	//if n <= 0 {
	//	return 0
	//}
	//res := 0
	//// 记录每个柱子 左边柱子的最大高度
	//leftMax := make([]int, n)
	//leftMax[0] = height[0]
	//// 第1个柱子不存雨水 故从索引1开始 向后/右 遍历
	//for i := 1; i < n; i++ {
	//	leftMax[i] = max(leftMax[i-1], height[i])
	//}
	//
	//// 记录每个柱子 右边柱子的最大高度
	//rightMax := make([]int, n)
	//rightMax[n-1] = height[n-1]
	//// 右边最后1个柱子不存雨水 故从n-2开始 向前/左 遍历
	//for i := n - 2; i >= 0; i-- {
	//	rightMax[i] = max(rightMax[i+1], height[i])
	//}
	//// 遍历数组 求和
	//for i, h := range height {
	//	res += min(leftMax[i], rightMax[i]) - h
	//}
	//return res

	// 单调栈解法
	res := 0
	stack := []int{}
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			res += curWidth * curHeight
		}
		stack = append(stack, i)
		fmt.Println(stack)
	}
	return res
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
}
