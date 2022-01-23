## 栈和队列
输入：[2,4,-2,-4,3,1,5], k=4
输出：[4,4,3,5]
即在k个窗口范围内找出最大值，如下
[2,4,-2,-4],3,1,5  -> 当前窗口最大值为 4
2,[4,-2,-4,3],1,5  -> 当前窗口最大值为 4
2,4,[-2,-4,3,1],5  -> 当前窗口最大值为 3
2,4,-2,[-4,3,1,5]  -> 当前窗口最大值为 3

package main

import (
	"fmt"
)

type myQueue struct {
	queue []int
}

func NewMyQueue() *myQueue { // 结构体的构造函数
	return &myQueue{ // 初始化
		queue: make([]int, 0),
	}
}

func (m *myQueue) front() int { // 返回队头元素
	return m.queue[0]
}
func (m *myQueue) back() int { // 返回队尾元素
	return m.queue[len(m.queue)-1]
}
func (m *myQueue) empty() bool { // 队列是否为空
	return len(m.queue) == 0
}
func (m *myQueue) push(val int) { // 入队列
	// 队列不为空 且 当前元素 大于 队尾元素
	for !m.empty() && val > m.back() {
		// 就把队尾元素去掉
		m.queue = m.queue[:len(m.queue)-1] // 即队列长度-1
	}
	// 上述循环完成后 队头元素 一定是最大的值
	m.queue = append(m.queue, val) // 把当前元素追加到 队尾
}
func (m *myQueue) pop(val int) { // 弹出元素 即队头元素
	// 如队列不为空 且 当前元素 等于 队头元素
	if !m.empty() && val == m.front() {
		m.queue = m.queue[1:] // 移除第1个元素 即队头元素
	}
}

// 暴力解法 可能会超时
func maxSlidingWindow(nums []int, k int) []int {
	//// 求切片中的最大值
	//var maxValOfArr func([]int) int
	//maxValOfArr = func(arr []int) int {
	//	max := arr[0]
	//	if len(arr) == 1 {
	//		return max
	//	}
	//	for i := 1; i < len(arr); i++ {
	//		if arr[i] > max {
	//			max = arr[i]
	//		}
	//	}
	//	return max
	//}
	//
	//res := []int{}
	//for i := 0; i+k <= len(nums); i++ {
	//
	//	res = append(res, maxValOfArr(nums[i:i+k]))
	//}
	//return res

	// 使用队列实现
	queue := NewMyQueue()
	n := len(nums)
	res := make([]int, 0)
	// 1. 先把 前k个 元素入队列
	for i := 0; i < k; i++ {
		queue.push(nums[i])
	}
	// 上述完成后 前k个 元素的最大值 一定在队头
	// 2. 再把 前k个 元素的最大值 追加到结果集
	res = append(res, queue.front()) // 即把队头元素 添加到结果集
	// 3. 再遍历k后面元素，即从k开始遍历
	for i := k; i < n; i++ {
		queue.pop(nums[i-k]) // 滑动窗口 移除 数组 最前面的元素
		queue.push(nums[i])  // 滑动窗口 把当前对应元素[即nums[i]] 入队列
		// 上述循环一次 都会更新队头元素 即当前窗台的最大值
		res = append(res, queue.front()) // 然后把 当前窗口的最大值 添加到结果集
	}
	return res
}

func main() {
	n := []int{2, 4, -2, -4, 3, 1, 5}
	fmt.Println(maxSlidingWindow(n, 4)) // [4 4 3 5]
}
