## 二叉树的属性
输入：root = [1,2,2,3,4,4,3]
输出：true
即：        1
        2       2
      3   4   4   3   
二叉的左右子树是对称的

// 递归
func isSymmetric(root *TreeNode) bool {
	// 把这个数设置成两个树，对比这两个树
	return isTree(root, root)
}
func isTree(a, b *TreeNode) bool {
	// 判断是否是空
	if a == nil && b == nil {
		return true
	}
	// 判断其中一个是否是空
	if a == nil || b == nil {
		return false
	}
	// 判断a.R 是否等于b.L
	return a.Val == b.Val && isTree(a.Right, b.Left) && isTree(a.Left, b.Right)
}

// 迭代
func isSymmetric(root *TreeNode) bool {
    var queue []*TreeNode;
    if root != nil {
        queue = append(queue, root.Left, root.Right);
    }
    for len(queue) > 0 {
        left := queue[0];
        right := queue[1];
        queue = queue[2:];
        if left == nil && right == nil {
            continue;
        }
        if left == nil || right == nil || left.Val != right.Val {
            return false;
        };
        queue = append(queue, left.Left, right.Right, right.Left, left.Right);
    }
    return true;
}
