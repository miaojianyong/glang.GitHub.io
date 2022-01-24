## 二叉树的修改与构造
一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：

二叉树的根是数组 nums 中的最大元素。
左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
返回有给定数组 nums 构建的 最大二叉树 

输入：nums = [3,2,1,6,0,5]
输出：[6,3,5,null,2,0,null,null,1]
即，如下二叉树：
      6
   /    \
  3      5
   \    / 
    2  0 
     \
      1
即根节点为数组中的最大值[6] 
  左子树子树是最大值左边部分[3,2,1]
  右子树子树是最大值右边部分[0,5]
且左右子树的根节点也是当前部分的最大值 即如上所示

func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums)<1{return nil}
    //首选找到最大值
    index:=findMax(nums)
    //其次构造二叉树
    root:=&TreeNode{
        Val: nums[index],
        Left:constructMaximumBinaryTree(nums[:index]),//左半边
        Right:constructMaximumBinaryTree(nums[index+1:]),//右半边
        }
    return root
}
func findMax(nums []int) (index int){
    for i:=0;i<len(nums);i++{
        if nums[i]>nums[index]{
            index=i
        }
    }
    return 
}
