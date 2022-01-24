## 二叉搜索树的属性
二叉搜索树（BST）的根节点和一个目标值值。 
需要在BST中找到节点值等于给定值的节点。 
返回以该节点为根的子树。 如果节点不存在，则返回 NULL
如下二叉搜索树:
        4
       / \
      2   7
     / \
    1   3 ----> 和目标值2
返回如下子树：
      2     
     / \   
    1   3  即返回目标值对应的节点为根的子树

// 递归
func searchBST(root *TreeNode, val int) *TreeNode {
	// 如果根节点为空 或 根节点的值就是要找的值 就返回根节点
	if root == nil || root.Val == val {
		return root
	}
	// 如果根节点的值 大于 要找的值 就去搜索左子树 因为左子树上的值都比根节点的值小
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	// 否则就去搜索右子树
	return searchBST(root.Right, val)
}

// 迭代法
func searchBST2(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val > val { // 如根节点的值大于val 去搜索左子树
			root = root.Left
		} else if root.Val < val { // 如根节点的值小于val 去搜索右子树
			root = root.Right
		} else { // 都不是就退出循环
			break
		}
	}
	return root
}
