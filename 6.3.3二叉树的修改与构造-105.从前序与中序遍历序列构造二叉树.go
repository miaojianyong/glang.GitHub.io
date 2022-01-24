## 二叉树的修改与构造
preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出：
     3
   /   \
  9    20
       / \
      15  7
即根据前序遍历出的结果 和 中序遍历出的结果 构造二叉树

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder)<1||len(inorder)<1{return nil}
    left:=findRootIndex(preorder[0],inorder)
    root:=&TreeNode{
        Val: preorder[0],
        Left: buildTree(preorder[1:left+1],inorder[:left]),
        Right: buildTree(preorder[left+1:],inorder[left+1:])}
    return root
}
func findRootIndex(target int,inorder []int) int{
    for i:=0;i<len(inorder);i++{
        if target==inorder[i]{
            return i
        }
    }
    return -1
}
