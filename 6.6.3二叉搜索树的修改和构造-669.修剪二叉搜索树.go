## 二叉搜索树的修改和构造
给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。
示例1：
输入：root = [1,0,2], low = 1, high = 2
输出：[1,null,2]
即：
       1
     /   \
    0     2 -->修剪成边界在[1,2]之间
即结果如下：
    1
     \
      2
示例2：
输入：root = [3,0,4,null,2,null,null,1], low = 1, high = 3
输出：[3,2,null,1]
即：
      3
     /  \
    0    4
     \
      2
     /
    1  -->修剪成边界在[1,3]之间
即结果如下：
        3
       /
      2
     /
    1

// 递归
func trimBST(root *TreeNode, low, high int) *TreeNode {
	// 如果是空数 返回空
	if root == nil {
		return nil
	}
	// 找符号 [low,high]区间的节点
	if root.Val < low { // 小于最小值 去右子树
		right := trimBST(root.Right, low, high)
		return right
	} else if root.Val > high { // 大于最大值 去左子树
		left := trimBST(root.Left, low, high)
		return left
	}
	// 给左右子树 赋值
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

// 迭代法
func trimBST2(root *TreeNode, low, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 处理root 让root移动到 [low,high]区间内 （左右都含）
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	cur := root
	// 此时root 已在[low,high]范围内 处理左孩子元素小于low的情况
	for cur != nil {
		for cur.Left != nil && cur.Left.Val < low {
			cur.Left = cur.Left.Right
		}
		cur = cur.Left
	}
	cur = root
	// 此时root 已在[low,high]范围内 处理右孩子元素大于high的情况
	for cur != nil {
		for cur.Right != nil && cur.Right.Val > high {
			cur.Right = cur.Right.Left
		}
		cur = cur.Right
	}
	return root
}
