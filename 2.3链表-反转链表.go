## 链表
描述：反转一个单链表。
示例: 输入: 1->2->3->4->5->NULL 输出: 5->4->3->2->1->NULL

//双指针
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode // 空指针
	cur := head       // 当前元素
	for cur != nil {
		next := cur.Next // 保存当前元素的下个节点
		cur.Next = pre   // 反转操作 如1->2变为1->nil、2->3变为2->1
		pre = cur        // 空指针指向当前元素
		cur = next       // 当前元素后移 当等于nil时结束循环
	}
	return pre
}

//递归
func reverseList2(head *ListNode) *ListNode {
	var reverse func(pre, head *ListNode) *ListNode
	reverse = func(pre, head *ListNode) *ListNode {
		if head == nil {
			return pre
		}
		next := head.Next          // 保存下个节点
		head.Next = pre            // 反转操作
		return reverse(head, next) // 递归
	}
	return reverse(nil, head)
}
