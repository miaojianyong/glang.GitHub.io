## 二叉树属性
二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1

func isBalanced(root *TreeNode) bool {
    if root==nil{
        return true
    }
    if !isBalanced(root.Left) || !isBalanced(root.Right){
        return false
    }
    LeftH:=maxdepth(root.Left)+1
    RightH:=maxdepth(root.Right)+1
    if abs(LeftH-RightH)>1{
        return false
    }
    return true
}
func maxdepth(root *TreeNode)int{
    if root==nil{
        return 0
    }
    return max(maxdepth(root.Left),maxdepth(root.Right))+1
}
func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
func abs(a int)int{
    if a<0{
        return -a
    }
    return a 
}
