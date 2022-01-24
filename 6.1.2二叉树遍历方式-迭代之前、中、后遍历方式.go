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

// preorderTraversal 前序遍历
func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	st := list.New()   // 定义链表
	st.PushBack(root)  // 把根节点加入链表
	for st.Len() > 0 { // 链表不为空再循环
		node := st.Remove(st.Back()).(*TreeNode) // 移除链表的元素 并返回
		res = append(res, node.Val)              //处理根节点
		if node.Right != nil {                   // 处理右孩子
			st.PushBack(node.Right)
		}
		if node.Left != nil { // 处理右孩子
			st.PushBack(node.Left)
		}
	}
	return res
}

// 中序遍历 左中右
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	st := list.New()
	cur := root // 指针指向根节点
	// 如果指针不为空 或 队列不为空 就循环，
	// 因为元素弹出后队列可能是空，但还要继续
	// 指针到最后时，也是空 但队列一定不为空，要继续，故需下述俩个条件满足其一即可
	for cur != nil || st.Len() > 0 {
		if cur != nil { // 如果指针不为空
			st.PushBack(cur) // 把当前指向的元素添加到队列
			// 左
			cur = cur.Left // 指针一直向下，走到 左面的最 底部位置

		} else { // 到底部后
			// 即按照出栈的方法 弹出栈顶的元素 即队列最后的元素 就是最左侧元素
			cur = st.Remove(st.Back()).(*TreeNode) // 把队列最后一个元素弹出
			// 中
			res = append(res, cur.Val) // 把弹出的元素 添加到结果集中
			// 右
			cur = cur.Right // 然后指针去出来右边的元素
		}
	}
	return res
}

// 后序遍历
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		res = append(res, node.Val) // 中
		// 把前序的 右->左 调整为 左->右
		if node.Left != nil {
			st.PushBack(node.Left)
		}
		if node.Right != nil {
			st.PushBack(node.Right)
		}
	}
	// 调用反转数组方法
	reverse(res)
	return res
}

// 反转数组
func reverse(a []int) {
	l, r := 0, len(a)-1 // 指向数组的第1个元素 和 最后1个元素
	for l < r {
		a[l], a[r] = a[r], a[l] // 交换
		l, r = l+1, r-1         // l前移1位 r后移1位
	}
}

func main() {
	root := &TreeNode{
		Val: 5,
	}
	root_l := &TreeNode{
		Val: 4,
	}
	root_r := &TreeNode{
		Val: 6,
	}
	root_l_l := &TreeNode{
		Val: 1,
	}
	root_l_r := &TreeNode{
		Val: 2,
	}
	root_r_l := &TreeNode{
		Val: 7,
	}
	root_r_r := &TreeNode{
		Val: 8,
	}
	root.Left = root_l
	root_l.Left = root_l_l
	root_l.Right = root_l_r
	root.Right = root_r
	root_r.Right = root_r_r
	root_r.Left = root_r_l
	fmt.Println(preorderTraversal(root))

	fmt.Println(inorderTraversal(root))

	fmt.Println(postorderTraversal(root))
}
