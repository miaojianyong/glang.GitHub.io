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
	  // check
    if root == nil {
        return root
    }
    // 相等 直接返回root节点即可
    if root == p || root == q {
        return root
    }
    // Divide
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)

    // Conquer
    // 左右两边都不为空，则根节点为祖先
    if left != nil && right != nil {
        return root
    }
    if left != nil {
        return left
    }
    if right != nil {
        return right
    }
    return nil
}
