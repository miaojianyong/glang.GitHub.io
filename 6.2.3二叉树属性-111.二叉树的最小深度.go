## 二叉树的属性
最小深度是从根节点到最近叶子节点的最短路径上的节点数量

// 递归
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 如果左子树为空 且右子树不为空 深度是右子树+1
	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	}
	// 如果右子树为空 且左子树不为空 深度是左子树+1
	if root.Right == nil && root.Left != nil {
		return minDepth(root.Left) + 1
	}
	// 左右子树都不为空 就是左右子树的最小值 +1
	leftDepth := minDepth(root.Left)        // 左
	rightDepth := minDepth(root.Right)      // 右
	depth := min(leftDepth, rightDepth) + 1 // 中
	return depth
	// 简写
	//return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

// 最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 迭代法
func minDepth2(root *TreeNode) int {
	levl := 0
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		levl++ // 进来说明有根节点 就+1
		for ; l > 0; l-- {
			node := queue[0]
			queue = queue[1:]
			// 如果左右子树都是空 就返回 说明到底了
			if node.Left == nil && node.Right == nil {
				return levl
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		l = len(queue)
	}
	return levl
}
