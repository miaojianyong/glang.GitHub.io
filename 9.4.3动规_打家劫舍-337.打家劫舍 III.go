## 动规_打家劫舍

小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
除了 root 之外，每栋房子有且只有一个“父“房子与之相连。
一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 

实例：
     3
    / \
   2   3
    \   \
     3   1
输入: root = [3,2,3,null,3,null,1]
输出: 7 
解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
即：根节点连着左右节点，就是只能获取这3个节点的1个，依次类推

// 动规1
func rob(root *TreeNode) int {
	res := robTree(root)
	return max(res[0], res[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robTree(cur *TreeNode) []int {
	if cur == nil {
		return []int{0, 0}
	}
	// 后序遍历
	left := robTree(cur.Left)
	right := robTree(cur.Right)

    // 考虑去偷当前的屋子
	robCur := cur.Val + left[0] + right[0]
    // 考虑不去偷当前的屋子
	notRobCur := max(left[0], left[1]) + max(right[0], right[1])

    // 注意顺序：0:不偷，1:去偷
	return []int{notRobCur, robCur}
}

// 动规-优化
func rob(root *TreeNode) int {
	var postOderTravel func (root *TreeNode)int
	postOderTravel = func(root *TreeNode) int {
		if root == nil{
			return 0
		}
		//l  r分别表示左右子树的最优解
		l := postOderTravel(root.Left)
		r := postOderTravel(root.Right)
		var temp int
		//root.Left.Val + root.Right.Val表示上上个的最优解
		if root.Left!=nil{
			temp = root.Left.Val
		}
		if root.Right!=nil{
			temp += root.Right.Val
		}
		//上上个的最优解 + 当前值
		temp += root.Val
		//last 表示子树（上一个）的最优解
		last := l+r
		// 使用Val来记录上一个的最优解
		root.Val = last
		return max(temp,last)
	}
	return postOderTravel(root)
}
