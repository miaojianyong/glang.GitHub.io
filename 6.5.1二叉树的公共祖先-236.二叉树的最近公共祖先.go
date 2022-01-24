## 二叉树的公共祖先
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
如：
      3
   /     \
  5       1
 / \     / \
6   2   0   8   
   / \
  7   4  --> 即节点5和1的公共根节点为3

// 递归
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 如当前节点为空 或 找到q或p节点 就返回当前节点
	if root == nil || root == p || root == q {
		return root
	}
	// 记录左右子树的返回值
	left := lowestCommonAncestor(root.Left, p, q)   // 左
	right := lowestCommonAncestor(root.Right, p, q) // 右
	// 处理左右子树的返回值
	// 如果左右子树的返回值都不为空 就返回当前节点
	if left != nil && right != nil {
		return root
	}
	if left != nil { // 如左子树不为空 就返回左子树
		return left
	}
	if right != nil { // 如右子树不为空 就返回左子树
		return right
	}
	return nil // 否则 返回空
}
