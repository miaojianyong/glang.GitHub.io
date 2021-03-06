## 二叉树的属性
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数

func max (a, b int) int {
    if a > b {
        return a;
    }
    return b;
}
// 递归
func maxDepth(root *treenode) int {
    if root == nil {
        return 0;
    }
    leftDepth := maxDepth(root.Left)        // 左
	rightDepth := maxDepth(root.Right)      // 右
	depth := max(leftDepth, rightDepth) + 1 // 中
	return depth
	// 简写
	//return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
// 迭代
func maxdepth(root *treenode) int {
    levl := 0;
    queue := make([]*treenode, 0);
    if root != nil {
        queue = append(queue, root);
    }
    for l := len(queue); l > 0; {
        for ;l > 0;l-- {
            node := queue[0];
            if node.left != nil {
                queue = append(queue, node.left);
            }
            if node.right != nil {
                queue = append(queue, node.right);
            }
            queue = queue[1:];
        }
        levl++;
        l = len(queue);
    }
    return levl;
}
