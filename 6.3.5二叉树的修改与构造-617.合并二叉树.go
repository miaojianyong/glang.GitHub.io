## 二叉树的修改与构造
即让两个二叉树合并成一个二叉树
输入：root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
输出：[3,4,5,5,4,null,7]
如：
root1:
     1
   /   \
  3     2
 /        
5            
root2:
     2
   /   \
  1     3
   \     \
    4     7
输出结果：
     3
   /   \
  4     5
 / \     \
5   4     7
即相同节点上的值相加，不同节点合并

 //前序遍历（递归遍历，跟105 106差不多的思路）
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    var value int 
    var nullNode *TreeNode//空node，便于遍历
    nullNode=&TreeNode{
        Val:0,
        Left:nil,
        Right:nil}
    switch {
        case t1==nil&&t2==nil: return nil//终止条件
        default : //如果其中一个节点为空，则将该节点置为nullNode，方便下次遍历
        if t1==nil{
            value=t2.Val
            t1=nullNode
        }else if t2==nil{
            value=t1.Val
            t2=nullNode
        }else {
            value=t1.Val+t2.Val
        }
    }
    root:=&TreeNode{//构造新的二叉树
        Val: value,
        Left: mergeTrees(t1.Left,t2.Left),
        Right: mergeTrees(t1.Right,t2.Right)}
    return root
}

// 前序遍历简洁版
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil {
        return root2
    }
    if root2 == nil {
        return root1
    }
    root1.Val += root2.Val
    root1.Left = mergeTrees(root1.Left, root2.Left)
    root1.Right = mergeTrees(root1.Right, root2.Right)
    return root1
}

// 迭代版本
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    queue := make([]*TreeNode,0)
    if root1 == nil{
        return root2
    }
    if root2 == nil{
        return root1
    }
    queue = append(queue,root1)
    queue = append(queue,root2)

    for size:=len(queue);size>0;size=len(queue){
        node1 := queue[0]
        queue = queue[1:]
        node2 := queue[0]
        queue = queue[1:]
        node1.Val += node2.Val
        // 左子树都不为空
        if node1.Left != nil && node2.Left != nil{
            queue = append(queue,node1.Left)
            queue = append(queue,node2.Left)
        }
        // 右子树都不为空
        if node1.Right !=nil && node2.Right !=nil{
            queue = append(queue,node1.Right)
            queue = append(queue,node2.Right)
        }
        // 树 1 的左子树为 nil，直接接上树 2 的左子树
        if node1.Left == nil{
            node1.Left = node2.Left
        }
        // 树 1 的右子树为 nil，直接接上树 2 的右子树
        if node1.Right == nil{
            node1.Right = node2.Right
        }
    }
    return root1
}
