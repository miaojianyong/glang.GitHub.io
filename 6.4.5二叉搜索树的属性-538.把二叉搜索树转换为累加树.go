## 二叉搜索树的属性
给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
如：[2, 5, 13]
第1个元素=`2+5+13`=20
第2个元素=`5+13`=18
第3个元素，就自己故=13
新数组为`[20,18,13]`

func convertBST(root *TreeNode) *TreeNode {
     var sum int
     RightMLeft(root,&sum)
     return root
}
func RightMLeft(root *TreeNode,sum *int) *TreeNode {
    if root==nil{return nil}//终止条件，遇到空节点就返回
    RightMLeft(root.Right,sum)//先遍历右边
    temp:=*sum//暂存总和值
    *sum+=root.Val//将总和值变更
    root.Val+=temp//更新节点值
    RightMLeft(root.Left,sum)//遍历左节点
    return root
}
