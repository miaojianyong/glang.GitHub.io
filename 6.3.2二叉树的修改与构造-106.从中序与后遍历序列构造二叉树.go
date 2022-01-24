## 二叉树的修改与构造
中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7

func buildTree(inorder, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	// 先找到根节点 即后序数组中的最后一个元素
	// 如：中[8,4,15,12,7] 后[8,15,7,12,4]
	rootValue := postorder[len(postorder)-1] // 即[4]
	// 按照上述根节点 分隔中序数组 即左边是左子树 右边是右子树
	rootIndex := findRootIndex(inorder, rootValue) // 即:1
	// 构造树
	root := &TreeNode{
		Val: rootValue, // [4]
		// 左子树 中序数组不含上述索引对应的值[:1]=8 后序数组不含上述索引对应的值[:1]=8
		Left: buildTree(inorder[:rootIndex], postorder[:rootIndex]),
		// 右子树 中序数组去掉上述索引对应的值再到最后[1+1:]=[15,12,7] 后序数组当前索引对应的值再到最后（不含数组最后一个值）[1:4]=[15,7,12]
		Right: buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1]),
	}
	return root
}

// 找到分隔中序数组的索引
func findRootIndex(inorder []int, target int) (index int) {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == target {
			return i
		}
	}
	return -1
}
