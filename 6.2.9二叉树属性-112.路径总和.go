## 二叉树的属性
指定一条路个节点的值 的和 等于目标值

// 递归
func isPathSum(root *TreeNode, targetSum int) bool {
	var flag bool // 是否找到标识
	if root == nil {
		return false
	}
	pathSum(root, 0, targetSum, &flag)
	return flag
}

func pathSum(root *TreeNode, sum, targetSum int, flag *bool) {
	sum += root.Val // 加上中间节点的值
	// 如果到叶子节点 且 sum=targetSum 表示找到了
	if root.Left == nil && root.Right == nil && sum == targetSum {
		*flag = true // 找到了 改变标识符值
	}
	// && !(*flag)好像去掉也行
	if root.Left != nil && !(*flag) { // 左节点不为空 且没找到
		pathSum(root.Left, sum, targetSum, flag)
	}
	if root.Right != nil && !(*flag) { // 右节点不为空 且没找到
		pathSum(root.Right, sum, targetSum, flag)
	}
}
