## 二叉树的属性
输入：root = [1,2,2,3,4,4,3]
输出：true
即：        1
        2       2
      3   4   4   3   
二叉的左右子树是对称的

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return defs(root.Left, root.Right)
}
func defs(left, right *TreeNode) bool {
	if left == nil && right == nil { // 左右都为空
		return true
	}
	if left == nil || right == nil { // 左或右 有1个为空
		return false
	}
	if left.Val != right.Val { // 左子树值 不等于 右子树值
		return false
	}
	var inside, outside bool
	inside = defs(left.Right, right.Left)  // 内侧 左子树：右 右子树：左
	outside = defs(left.Left, right.Right) // 外侧 左子树：左 右子树：右
	isSame := inside && outside            // 内侧 外侧是否都对称 左右子树：中
	return isSame
	// 上述三行简写
	//return defs(left.Right,right.Left)&&defs(left.Left,right.Right)
}

// 迭代
func isSymmetric2(root *TreeNode) bool {
	var queue []*TreeNode
	if root != nil { // 把左右子树追加到切片
		queue = append(queue, root.Left, root.Right)
	}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue // 表示到底了 继续下一次比较
		}
		// 如果左右子树有1个为空 或对应的值不同
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		// 把左右子树的外侧 内侧追加到切片
		// left.Left, right.Right 外侧
		// left.Right, right.Left 内侧
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}
