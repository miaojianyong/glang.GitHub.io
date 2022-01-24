## 二叉搜索树的属性
一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效 二叉搜索树定义如下：
节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
如
     2
   /   \
  1     3 ---> 返回 true
即根节点的值如2 大于 它的左节点的值如1
根节点的值如2 小于 它的右节点的值如3

// 递归 即递归的验证左右子节点的值
func isValidBST(root *TreeNode) bool {
	// 空数也是二叉搜索树
	if root == nil {
		return true
	}
	// 限制数值范围
	return check(root, math.MinInt64, math.MaxInt64)
}

func check(node *TreeNode, min, max int64) bool {
	if node == nil { // 到叶子节点时返回true
		return true
	}
	// 如果当前值超出范围 返回false
	if int64(node.Val) <= min || int64(node.Val) >= max {
		return false
	}
	// 分别判断左右子树 都符合才返回true
	// 左子树 最小值就是min 最大值就是当前值
	// 右子树 最小值就是当前值 最大值是max
	return check(node.Left, min, int64(node.Val)) && check(node.Right, int64(node.Val), max)
}

// 中序遍历 即根据二叉搜索的特性[二叉搜索树的中序遍历结果是升序]
func isValidBST2(root *TreeNode) bool {
	var prev *TreeNode // 用于保存前一个节点
	var travel func(node *TreeNode) bool
	travel = func(node *TreeNode) bool {
		if node == nil { // 空树也是二叉搜索树
			return true
		}
		leftRes := travel(node.Left) // 左
		// 如当前值小于等于前一个节点的值 返回false 因为中序遍历是升序
		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node                    // 保存前一个节点
		rightRes := travel(node.Right) // 右
		return leftRes && rightRes
	}
	return travel(root)
}
