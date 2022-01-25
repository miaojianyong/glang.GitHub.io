## 回溯_组合
找出所有相加之和为 n 的 k 个数的组合。
组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字

输入: k = 3, n = 7
输出: [[1,2,4]]

func combineSum3(n, k int) [][]int {
	var res [][]int
	var tack []int
	// 函数中修改 切片的值 传递指针
	var backTree func(n, k, startIndex int, tack *[]int, res *[][]int)
	backTree = func(n, k, startIndex int, tack *[]int, res *[][]int) {
		if len(*tack) == k {
			var sum int
			temp := make([]int, k)
			for k, v := range *tack {
				sum += v
				temp[k] = v // 存储索引和对应的值
			}
			if sum == n { // 把符合条件的组合存储到结果集
				*res = append(*res, temp)
			}
			return
		}
		//(k-len(*tack)) 表示还剩多少个可填充的元素
		for i := startIndex; i <= 9-(k-len(*tack))+1; i++ {
			*tack = append(*tack, i)       // 记录路径
			backTree(n, k, i+1, tack, res) // 递归
			*tack = (*tack)[:len(*tack)-1] // 回溯
		}
	}
	backTree(n, k, 1, &tack, &res)
	return res
}

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var path []int
	var backTrack func(int, int)
	backTrack = func(sum, start int) {
		// 剪枝，如果当前路径和已经大于n,后续遍历就没有意义了
		if sum > n {
			return
		}
		// 递归终止条件
		// 如果当前路径长度为k,且路径和等于n，便找到了一条满足要求的路径
		if len(path) == k && sum == n {
			temp := make([]int, k)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		// for循环遍历(剪枝)
		for i := start; i <= 9-(k-len(path))+1; i++ {
			// 处理每一个节点
			sum += i
			path = append(path, i)
			// 递归
			backTrack(sum, i+1)
			// 回溯
			sum -= i
			path = path[:len(path)-1]
		}
	}
	backTrack(9, 1)
	return res
}

func main() {
	fmt.Println(combineSum3(9, 3)) // [[1 2 6] [1 3 5] [2 3 4]]
	fmt.Println(combinationSum3(9, 3))
}
