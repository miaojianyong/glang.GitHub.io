## 链表
描述：删除链表的倒数第n个节点
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

// 双指针
func removeNthFromEnd(head *ListNode19, n int) *ListNode19 {
	dummyHead := &ListNode19{} // 定义虚拟头节点
	dummyHead.Next = head      // 让其指向头节点
	cur := head                // 定义当前节点指向头节点
	prev := dummyHead          // 指针 指向虚拟节点
	i := 1                     // 从1开始
	for cur != nil {           // 循环让cur指向下个节点 当执行空时退出
		cur = cur.Next
		if i > n { // 当i大于n时 prev才指向下个节点
			prev = prev.Next
		}
		i++
	}
	// 上述循环结束时 prev会在要删除的节点前一个位置
	// 如1->2->3-4 n=2, prev会在节点2的位置
	// 让prev的下个节点指向 下下个节点 即删除了节点3
	prev.Next = prev.Next.Next
	return dummyHead.Next
}
