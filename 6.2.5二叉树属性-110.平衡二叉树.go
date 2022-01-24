## 二叉树属性
二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1

// 递归
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 判断左右子树是否是平衡二叉树
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	// 左右子树的高度
	leftH := maxDepth(root.Left) + 1
	rightH := maxDepth(root.Right) + 1
	// 判断左右子树差是否大于1 绝对值
	if abs(leftH-rightH) > 1 {
		return false
	}
	return true
}

// 左右子树的高度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 绝对值
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
