## 二叉搜索树的修改和构造
一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。

示例：
输入：nums = [1,3]
输出：[3,1]
解释：[1,3] 和 [3,1] 都是高度平衡二叉搜索树。
即：
    3       1
   /         \
  1    和     3 -->这两种数都符合要求

// 递归 前序
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// 按照二叉树特点 数组的中间元素 就是跟节点
	root := &TreeNode{Val: nums[len(nums)/2]}
	// 数组的左边 为左子树
	root.Left = sortedArrayToBST(nums[:len(nums)/2])
	// 数组的右边 为右子树
	root.Right = sortedArrayToBST(nums[len(nums)/2+1:])
	return root
}
