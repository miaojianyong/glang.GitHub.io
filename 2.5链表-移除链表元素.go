## 链表
描述：删除链表中等于给定值 val 的所有节点。
输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]
// 方式1 递归
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil { // 终止条件
		return head
	}
	head.Next = removeElements(head.Next, val) // 递归
	if head.Val == val {                       // 如果当前元素的值 等于 目标值
		return head.Next // 让当前元素 指向下个结点
	}
	return head
}

// 方式2 迭代
func removeElements2(head *ListNode, val int) *ListNode {
	// 处理头结点
	for head != nil && head.Val == val { // 如当前头结点不为空 且 值等于目标值
		head = head.Next // 让头结点指向下个结点
	}
	if head == nil {
		return head
	}
	pre := head           // 定义指针指向头结点
	for pre.Next != nil { // 用pre删除对应的目标元素
		// 如pre的下个结点的值 等于 目标值
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next // pre的下个结点 指向 下下个结点
		} else { // 否则 让pre执行下个结点 继续循环查询
			pre = pre.Next
		}
	}
	return head
}

// 双指针
func removeElements3(head *ListNode, val int) *ListNode {
	for head != nil { // 处理头结点
		if head.Val != val { // 如头结点的值 不等于 目标值
			break // 就退出
		}
		head = head.Next // 否则头结点指向 下个节点 即删除该节点
	}
	pre, cur := head, head
	for cur != nil {
		if cur.Val == val { // 如当前节点的值 等于 目标值
			pre.Next = cur.Next // 让pre指向 当前节点的下个结点
		} else { // 否则让pre和cur移动到相同节点
			pre = cur
		}
		cur = cur.Next // 让当前节点指向下个结点 继续循环查找
	}
	return head
}

// 虚拟头节点
func removeElements4(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{} // 创建虚拟节点
	dummyHead.Next = head    // 让虚拟节点指向头节点
	cur := dummyHead         // 让变量指向 虚拟节点
	// 当前节点不为空 且 下个结点不为空
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val { // 找到目标值
			cur.Next = cur.Next.Next // 让其指向 下下个节点
		} else { // 否则
			cur = cur.Next // 指向下个节点
		}
	}
	return dummyHead.Next // 返回虚拟节点后续的节点
}
