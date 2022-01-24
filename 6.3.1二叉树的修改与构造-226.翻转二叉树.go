## 二叉树的修改与构造
翻转一棵二叉树
输入：           
     4
   /   \
  2     7
 / \   / \
1   3 6   9   
输出：
     4
   /   \
  7     2
 / \   / \
9   6 3   1
即让二叉树的左右子树和对应的各个子节点交换位置

package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 反转二叉树
// 递归方式 -- 前序
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 先交换左右节点 再反转左右子树
	root.Left, root.Right = root.Right, root.Left // 类似 交换
	invertTree(root.Left)                         // 左
	invertTree(root.Right)                        // 右
	return root
}

// 递归方式 -- 后序
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)                         // 左
	invertTree(root.Right)                        // 右
	root.Left, root.Right = root.Right, root.Left // 类似 交换
	return root
}

// 迭代方式 -- 前序
func invertTree3(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			node.Left, node.Right = node.Right, node.Left // 交换
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	return root
}

// 迭代方式 -- 后序
func invertTree4(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	node := root
	var prev *TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 如果右节点是空 或 右节点等于指针变量 就交换
		if node.Right == nil || node.Right == prev {
			node.Left, node.Right = node.Right, node.Left
			prev = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
		}
	}
	return root
}

// 迭代方式 -- 层序遍历
func invertTree5(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := list.New()
	node := root
	queue.PushBack(node)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			e := queue.Remove(queue.Front()).(*TreeNode)
			e.Left, e.Right = e.Right, e.Left
			if e.Left != nil {
				queue.PushBack(e.Left)
			}
			if e.Right != nil {
				queue.PushBack(e.Right)
			}
		}
	}
	return root
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

	fmt.Println(invertTree(root))
}
