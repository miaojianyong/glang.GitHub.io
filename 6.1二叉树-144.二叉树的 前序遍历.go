## 二叉树的遍历方式

// 递归
func preorderTraversal(root *TreeNode) (res []int) {
    var traversal func(node *TreeNode)
    traversal = func(node *TreeNode) {
	if node == nil {
            return
	}
	res = append(res,node.Val) // 中
	traversal(node.Left) // 左
	traversal(node.Right) // 右
    }
    traversal(root)
    return res
}
// 迭代
func preorderTraversal(root *TreeNode) []int {
    ans := []int{}
	if root == nil {
		return ans
	}
	st := list.New()
    st.PushBack(root)
    for st.Len() > 0 {
        node := st.Remove(st.Back()).(*TreeNode)
        ans = append(ans, node.Val)
        if node.Right != nil {
            st.PushBack(node.Right)
        }
        if node.Left != nil {
            st.PushBack(node.Left)
        }
    }
    return ans
}
