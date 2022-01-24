## 二叉树的属性
在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置

//本题直接就是求有多少个节点，无脑存进数组算长度就行了。
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    res := 1 // 根节点
    if root.Right != nil { // 遍历右子树节点
        res += countNodes(root.Right)
    }
    if root.Left != nil { // 遍历左子树节点
        res += countNodes(root.Left)
    }
    return res
}
// 利用完全二叉树特性的递归解法
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftH, rightH := 0, 0
    leftNode := root.Left
    rightNode := root.Right
    for leftNode != nil {
        leftNode = leftNode.Left
        leftH++
    }
    for rightNode != nil {
        rightNode = rightNode.Right
        rightH++
    }
    if leftH == rightH {
        return (2 << leftH) - 1
    }
    return countNodes(root.Left) + countNodes(root.Right) + 1
}
