## 二叉树的属性
指定所有的路径和为目标值的集合

func pathsum(root *treenode, targetsum int) [][]int {
    var result [][]int//最终结果
    if root==nil{
        return result
    }
    var sumnodes []int//经过路径的节点集合
    haspathsum(root,&sumnodes,targetsum,&result)
    return result
}
func haspathsum(root *treenode,sumnodes *[]int,targetsum int,result *[][]int){
    *sumnodes=append(*sumnodes,root.val)
    if root.left==nil&&root.right==nil{//叶子节点
        fmt.println(*sumnodes)
        var sum int
        var number int
        for k,v:=range *sumnodes{//求该路径节点的和
            sum+=v
            number=k
        }
        tempnodes:=make([]int,number+1)//新的nodes接受指针里的值，防止最终指针里的值发生变动，导致最后的结果都是最后一个sumnodes的值
        for k,v:=range *sumnodes{
            tempnodes[k]=v
        }
        if sum==targetsum{
            *result=append(*result,tempnodes)
        }
    }
    if root.left!=nil{
        haspathsum(root.left,sumnodes,targetsum,result)
        *sumnodes=(*sumnodes)[:len(*sumnodes)-1]//回溯
    }
    if root.right!=nil{
        haspathsum(root.right,sumnodes,targetsum,result)
        *sumnodes=(*sumnodes)[:len(*sumnodes)-1]//回溯
    }
}
