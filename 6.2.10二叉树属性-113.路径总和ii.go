## 二叉树的属性
指定所有的路径和为目标值的集合

// 递归
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{} // 结果集
	path := []int{}  // 路径集
	var paths func(*TreeNode, int)
	paths = func(node *TreeNode, targetSum int) {
		if node == nil {
			return
		}
		targetSum -= node.Val         // 减中间节点的值
		path = append(path, node.Val) // 把中间节点追加到路径集
		defer func() {
			path = path[:len(path)-1]
		}()
		// 如到叶子节点 且目标和为0
		if node.Left == nil && node.Right == nil && targetSum == 0 {
			// 表示找到到该路径
			res = append(res, append([]int(nil), path...))
			return
		}
		paths(node.Left, targetSum)
		paths(node.Right, targetSum)
	}
	paths(root, targetSum)
	return res
}

func pathSum2(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	path := []int{}
	isPath(&res, root, targetSum, path)
	return res
}

// 要改变 二维切片数组 res := [][]int{} 需传递地址
func isPath(res *[][]int, root *TreeNode, targetSum int, path []int) {
	if root == nil {
		return
	}
	targetSum -= root.Val
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		cp := make([]int, len(path))
		copy(cp, path)          // 拷贝路径
		*res = append(*res, cp) // 把符合的路径追到到结果集
	}
	isPath(res, root.Left, targetSum, path)
	isPath(res, root.Right, targetSum, path)
	//targetSum += root.Val
	path = path[:len(path)-1] // 移除路径的最后一个元素，回溯
	// 如，5>4>1 到底了 遍历到了4的左孩子
	// 就执行上代码 进行回溯，即 5>4
	// 然后 去遍历 4的右孩子 即 5>4>2 依此逻辑遍历整颗树
}
