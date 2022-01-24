## 二叉树的属性
二叉树的左叶子值的和 包含左右子树的所有左叶子节点

// 递归
func sumOfLeftLeaves(root *TreeNode) int {
    var  res int
    findLeft(root,&res)
    return res
}
func findLeft(root *TreeNode,res *int){
    //左节点
    if root.Left!=nil&&root.Left.Left==nil&&root.Left.Right==nil{
        *res=*res+root.Left.Val
    }
    if root.Left!=nil{
        findLeft(root.Left,res)
    }
    if root.Right!=nil{
        findLeft(root.Right,res)
    }
}
// 迭代
func sumOfLeftLeaves(root *TreeNode) int {
    var  res int
    queue:=list.New()
    queue.PushBack(root)
    for queue.Len()>0{
        length:=queue.Len()
        for i:=0;i<length;i++{
            node:=queue.Remove(queue.Front()).(*TreeNode)
            if node.Left!=nil&&node.Left.Left==nil&&node.Left.Right==nil{
                res=res+node.Left.Val
            }
            if node.Left!=nil{
                queue.PushBack(node.Left)
            }
            if node.Right!=nil{
                queue.PushBack(node.Right)
            }
        }
    }
    return res
}
