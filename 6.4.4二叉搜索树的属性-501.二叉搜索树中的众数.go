## 二叉搜索树的属性
一个含重复值的二叉搜索树（BST）的根节点 root ，找出并返回 BST 中的所有 众数（即，出现频率最高的元素）
如：
    1
      \
       2
      /
     2 --> 输出2
即输出二叉搜索树中出现次数最高的元素[不止1个]

// 递归 普通二叉树 即中序遍历的结果不是升序的
func findMode(root *TreeNode) []int {
	var history map[int]int
	history = make(map[int]int)
	traversal(root, history)
	var maxCount int // 最大次数
	var maxIndex int // 对应的索引
	var res []int    // 结果集
	// 遍历map
	for i, v := range history {
		// 如果当前值的次数 大于最大次数
		if v > maxCount {
			maxCount = v // 把次数给 最大次数变量
			maxIndex = i // 当前元素的索引给
		}
	}
	for i, v := range history {
		// 如果当前数等于 map索引对应的数
		if v == history[maxIndex] {
			res = append(res, i) // 就把当前值的k 追加到结果集中
		}
	}
	return res
}
func traversal(root *TreeNode, history map[int]int) {
	if root.Left != nil { // 左
		traversal(root.Left, history)
	}
	// 判断当前值是否在map中
	if value, ok := history[root.Val]; ok {
		history[root.Val] = value + 1 // 如果在就次数+1
	} else {
		history[root.Val] = 1 // 否则就是第1次出现 次数为1
	}
	if root.Right != nil {
		traversal(root.Right, history)
	}
}

// 二叉搜索树 即中序的递增的
func findMode2(root *TreeNode) []int {
	res := make([]int, 0)
	count := 1
	max := 1
	var prev *TreeNode
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left) // 左
		// 如果前一个节点不为空 且 当前节点的值等于前一个节点的值
		if prev != nil && node.Val == prev.Val {
			count++ // 次数+1 即该值不是第1次出现
		} else {
			count = 1 // 否则 该值是第1次出现
		}
		// 如果次数 大于等于 max
		if count >= max {
			// 如果次数大于max 且 res结果集 不为空
			if count > max && len(res) > 0 {
				res = []int{node.Val} // 就把当前值 替换 到结果集中
			} else {
				res = append(res, node.Val) // 否则就把当前值 追加 到结果集中
			}
			max = count // 把次数给max
		}
		prev = node        // 记录前一个节点
		travel(node.Right) // 右
	}
	travel(root)
	return res
}
