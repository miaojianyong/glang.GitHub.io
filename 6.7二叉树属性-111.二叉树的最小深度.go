## 二叉树的属性
最小深度是从根节点到最近叶子节点的最短路径上的节点数量

func min(a, b int) int {
    if a < b {
        return a;
    }
    return b;
}
// 递归
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0;
    }
    if root.Left == nil && root.Right != nil {
        return 1 + minDepth(root.Right);
    }
    if root.Right == nil && root.Left != nil {
        return 1 + minDepth(root.Left);
    }
    return min(minDepth(root.Left), minDepth(root.Right)) + 1;
}

// 迭代
func minDepth(root *TreeNode) int {
    dep := 0;
    queue := make([]*TreeNode, 0);
    if root != nil {
        queue = append(queue, root);
    }
    for l := len(queue); l > 0; {
        dep++;
        for ; l > 0; l-- {
            node := queue[0];
            if node.Left == nil && node.Right == nil {
                return dep;
            }
            if node.Left != nil {
                queue = append(queue, node.Left);
            }
            if node.Right != nil {
                queue = append(queue, node.Right);
            }
            queue = queue[1:];
        }
        l = len(queue);
    } 
    return dep;
}
