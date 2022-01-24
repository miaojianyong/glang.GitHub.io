## 二叉搜索树的修改和构造
二叉搜索树（BST）的根节点 root 和要插入树中的值 value ，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点
输入：root = [4,2,7,1,3], val = 5
输出：[4,2,7,1,3,5]
如：
      4
    /   \
  2      7
 / \     
1   3 --> 加节点 5
输出：
      4
    /   \
  2      7
 / \    /
1   3  5
或者输出：
      5
    /   \
  2      7
 / \     
1   3
     \
      4 --> 该形式的二叉搜索树也满足要求 即加入的节点[符合二叉搜索树的模式就行]

// 递归
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 如果当前节点为空 就是要插入节点的位置 如赋值操作
	if root == nil {
		root = &TreeNode{ // 赋值
			Val: val,
		}
		return root // 返回
	}
	// 如果当前节点的值 大于 要插入的值 就去左子树
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else { // 否则 就去右子树
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// 迭代法
func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	node := root
	var pnode *TreeNode // 记录上一个节点
	for node != nil {
		if node.Val > val { // 当前节点的值大于 val就去左边
			pnode = node
			node = node.Left
		} else { // 否则就去 右边
			pnode = node
			node = node.Right
		}
	}
	// 用pnode赋值
	if pnode.Val > val {
		pnode.Left = &TreeNode{Val: val}
	} else {
		pnode.Right = &TreeNode{Val: val}
	}
	return root
}
