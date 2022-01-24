## 二叉搜索树的属性
二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 
如：
     1
   /   \
  0    48
       / \
     12   49 ---> 输出：1
即上述二叉树节点的值的最小差值为1[绝对值] 即 1-0=1/0-1=-1

// 递归
func getMinimumDifference(root *TreeNode) int {
	var res []int
	findMin(root, &res)
	min := 1000000 // 最大值的范围
	// 遍历数组
	for i := 1; i < len(res); i++ {
		// 让当前值 减去 上一个值
		tempVal := res[i] - res[i-1]
		if tempVal < min { // 如果 差值小于min 最小差就等于min
			min = tempVal
		}
	}
	return min
}
func findMin(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	findMin(root.Left, res)
	*res = append(*res, root.Val)
	findMin(root.Right, res)
}

// 在遍历的时就进行计算
func getMinimumDifference2(root *TreeNode) int {
	var prev *TreeNode // 保存上个节点
	min := 1000000
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left) // 左
		if prev != nil && node.Val-prev.Val < min {
			min = node.Val - prev.Val
		}
		prev = node // 保存上个节点
		travel(node.Right)
	}
	travel(root)
	return min
}
