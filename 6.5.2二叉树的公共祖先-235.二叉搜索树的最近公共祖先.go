## 二叉树的公共祖先
一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6 
解释: 节点 2 和节点 8 的最近公共祖先是 6。
如：
      6
   /     \
  2       8
 / \     / \
0   4   7   9   
   / \
  3   5  --> 即节点2 和节点 8 的最近公共祖先是 6，即区别于普通二叉树 该二叉搜索树每个节点的值排列不同

// 递归
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 如当前节点的值大于 p和q 就去左边找
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
		// 如当前节点的值小于 p和q 就去右边找
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else { // 否则当前节点的值 就在q和p之间 返回当前节点
		return root
	}
}
