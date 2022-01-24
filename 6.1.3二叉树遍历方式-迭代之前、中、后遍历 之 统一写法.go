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

// 前中后序 之 统一写法 迭代
// preorderTraversal 前序遍历 中、左、右
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := list.New()
	res := []int{}
	stack.PushBack(root)
	var node *TreeNode
	for stack.Len() > 0 {
		e := stack.Back()   // 返回链表最后1个元素
		stack.Remove(e)     // 弹出元素
		if e.Value == nil { // 如果为空，则表示要处理中间节点
			e = stack.Back()            // 弹出元素，即中间节点
			stack.Remove(e)             // 删除中间节点
			node = e.Value.(*TreeNode)  // 把该节点进行类型转换
			res = append(res, node.Val) // 把中间节点追加到结果集中
			continue                    // 继续弹出栈中下一个节点
		}
		node = e.Value.(*TreeNode) // 类型转换
		// 把弹出的元素 压入栈，顺序是：右、左、中，即前序遍历的倒序
		if node.Right != nil { // 右
			stack.PushBack(node.Right)
		}
		if node.Left != nil { // 左
			stack.PushBack(node.Left)
		}
		stack.PushBack(node) // 中
		stack.PushBack(nil)  // 中间节点压入栈后，在压入空节点作为中间节点的标记
	}
	return res
}

// 中序遍历 左中右
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return nil
	}
	st := list.New() // 栈
	st.PushBack(root)
	var node *TreeNode
	for st.Len() > 0 {
		e := st.Back()
		st.Remove(e)
		// 等于空 表示要处理中间节点
		if e.Value == nil {
			e = st.Back()               // 弹出中间节点
			st.Remove(e)                // 从栈中删除中间节点
			node = e.Value.(*TreeNode)  // 类型转换
			res = append(res, node.Val) // 把中间节点追加到结果集
			continue                    // 继续弹出栈中下一个节点
		}
		node = e.Value.(*TreeNode)
		// 入栈顺序 ：右中左，即中序遍历 左中右的倒序
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		st.PushBack(node) // 中间节点
		st.PushBack(nil)  // 压入空节点作为中间节点的标识
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return res
}

// 后序遍历
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := list.New()
	stack.PushBack(root)
	var node *TreeNode
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		if e.Value == nil {
			e = stack.Back()
			stack.Remove(e)
			node = e.Value.(*TreeNode) // 把弹出的元素进行类型转换然后赋值给node
			res = append(res, node.Val)
			continue
		}
		node = e.Value.(*TreeNode)
		// 入栈顺序：中右左 即遍历 左右中的倒序
		stack.PushBack(node)
		stack.PushBack(nil)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
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

	fmt.Println(inorderTraversal(root))

	fmt.Println(postorderTraversal(root))
}
