## 二叉树的属性
指定一条路个节点的值 的和 等于目标值

// 递归
func haspathsum(root *treenode, targetsum int) bool {
    var flage bool //找没找到的标志
    if root==nil{
        return flage
    }
    pathsum(root,0,targetsum,&flage)
    return flage
}
func pathsum(root *treenode, sum int,targetsum int,flage *bool){
    sum+=root.val
    if root.left==nil&&root.right==nil&&sum==targetsum{
        *flage=true
        return
    }
    if root.left!=nil&&!(*flage){//左节点不为空且还没找到
        pathsum(root.left,sum,targetsum,flage)
    } 
    if root.right!=nil&&!(*flage){//右节点不为空且没找到
        pathsum(root.right,sum,targetsum,flage)
    }
}
