package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preorderTraversal 前序遍历1
// 1. 确定参数和返回值 即root是参数 要得到树的各个节点的值，即用切片
func preorderTraversal(root *TreeNode) (res []int) {
	// 2. 确定终止条件 即当遍历到节点是空时，就退出
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil { // 终止条件
			return
		}
		// 3. 确定单层递归的逻辑
		// 前序 -> 中、左、右
		res = append(res, node.Val) // 中
		traversal(node.Left)        // 左
		traversal(node.Right)       // 右
	}
	traversal(root)
	return res
}

// 前序遍历2
func preorderTraversal2(root *TreeNode) (res []int) {
	// 2. 确定终止条件 即当遍历到节点是空时，就退出
	if root == nil { // 终止条件
		return
	}
	// 3. 确定单层递归的逻辑
	// 前序 -> 中、左、右
	//if root != nil {
	res = append(res, root.Val)                          // 中
	res = append(res, preorderTraversal2(root.Left)...)  // 左
	res = append(res, preorderTraversal2(root.Right)...) // 右
	//}
	return res
}

// inorderTraversal 中序遍历
func inorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)        // 左
		res = append(res, node.Val) // 中
		traversal(node.Right)       // 右
	}
	traversal(root)
	return res
}
func inorderTraversal2(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	res = append(res, inorderTraversal2(root.Left)...)  // 左
	res = append(res, root.Val)                         // 中
	res = append(res, inorderTraversal2(root.Right)...) // 右
	return res
}

// postorderTraversal 后序遍历
func postorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)        // 左
		traversal(node.Right)       // 右
		res = append(res, node.Val) // 中
	}
	traversal(root)
	return res
}
func postorderTraversal2(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	res = append(res, postorderTraversal2(root.Left)...)  // 左
	res = append(res, postorderTraversal2(root.Right)...) // 右
	res = append(res, root.Val)                           // 中
	return res
}

func main() {
	root := &TreeNode{
		Val: 1,
	}
	root_l := &TreeNode{
		Val: 2,
	}
	root_r := &TreeNode{
		Val: 3,
	}
	root_l_l := &TreeNode{
		Val: 4,
	}
	root_r_r := &TreeNode{
		Val: 5,
	}
	root.Left = root_l
	root_l.Left = root_l_l
	root.Right = root_r
	root_r.Right = root_r_r
	fmt.Println(preorderTraversal(root))
	fmt.Println(preorderTraversal2(root))

	fmt.Println(inorderTraversal(root))
	fmt.Println(inorderTraversal2(root))

	fmt.Println(postorderTraversal(root))
	fmt.Println(postorderTraversal2(root))
}
