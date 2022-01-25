## 回溯-组合
两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合
输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
即输出 1-n范围内的所有k个元素的组合

func combine(n, k int) [][]int {
	var res [][]int
	// 如果n集合小于等于0 k组合小于等于0 k组合大于n集合
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	var backtrack func(n, k, start int, track []int)
	backtrack = func(n, k, start int, track []int) {
		// 如track中存储了 k个元素
		if len(track) == k {
			temp := make([]int, k)
			copy(temp, track)       // 从track中拷贝结果 即一次组合
			res = append(res, temp) // 把拷贝的组合追加到结果集
		}
		//if len(track)+n-start+1 < k { // 如不能构成k组合，就退出 不写也行
		//	return
		//}
		for i := start; i <= n; i++ { // 如 1. 循环1-4这个集合 从1开始
			// 2. 把1追加到track
			track = append(track, i) // 从集合取出1个元素 追加到track
			// 3. 递归 即 backtrack(4,2,2,[1])
			// 4. 即继续把 2 追加到track=[1,2] -- 此时符合上述 len(track) == k 会把[1,2]追加到结果集
			backtrack(n, k, i+1, track) // 递归 去看看track中的元素是否符合标准
			// 回溯 即[1,2] --> [2],保留最后一个元素 继续循环
			track = track[:len(track)-1] // 回溯 即仅存最后一个元素
		}
	}
	backtrack(n, k, 1, []int{})
	return res
}

// 优化
func combine2(n int, k int) [][]int {
	res := [][]int{}
	path := []int{}
	var backtrack func(int)
	backtrack = func(first int) {
		if len(path) == k {
			res = append(res, append([]int(nil), path...))
			return
		}
		for i := first; n-i+1 >= k-len(path); i++ {
			path = append(path, i)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(1)
	return res
}
