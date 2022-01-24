## 二叉树的属性
二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

// 递归 前序 中、左、右
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	// 根节点 和路径
	var travel func(node *TreeNode, s string)
	travel = func(node *TreeNode, s string) {
		// 当当前节点的没有左右孩子时 即表示到了叶子节点
		if node.Left == nil && node.Right == nil {
			// 把参数s(即根节点到该节点的路径)和当前节点的值进行拼接
			path := s + strconv.Itoa(node.Val)
			// 上述path就是一条路径，把该路径追加到结果集中
			res = append(res, path)
			return
		}
		// 处理中间节点
		s = s + strconv.Itoa(node.Val) + "->" // 中
		// 当前节点有左孩子就继续递归
		if node.Left != nil { // 左
			travel(node.Left, s)
		}
		// 当前节点有右孩子就继续递归
		if node.Right != nil { // 右
			travel(node.Right, s)
		}
	}
	travel(root, "")
	return res
}

// 迭代法
func binaryTreePaths2(root *TreeNode) []string {
	stack := []*TreeNode{}     // 保存遍历该数的节点
	paths := make([]string, 0) // 保存每条路径
	res := make([]string, 0)   // 结果集
	if root != nil {
		stack = append(stack, root)
		paths = append(paths, "")
	}
	for len(stack) > 0 {
		l := len(stack)
		// 取出当前节点 并清除该节点 类似栈的弹出和移除操作
		node := stack[l-1]
		path := paths[l-1]
		stack = stack[:l-1]
		paths = paths[:l-1]
		// 如果是叶节点 就把当前节点拼接遍历的路径 然后追加到结果集
		if node.Left == nil && node.Right == nil {
			res = append(res, path+strconv.Itoa(node.Val))
			continue // 继续下一条路径
		}
		// 如果当前节点的左孩子不为空
		if node.Left != nil {
			// 就把当前节点追加到stack
			stack = append(stack, node.Left)
			// 然后拼接路径 然后追加到paths
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
		// 如果当前节点的右孩子不为空 如上操作
		if node.Right != nil {
			stack = append(stack, node.Right)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
	}
	return res
}
