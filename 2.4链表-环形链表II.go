## 链表
描述：一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

func detectCycle(head *ListNode) *ListNode {
    // 哈希表
    // 即遍历链表元素 并 存入map中，如能在哈希表中取到之前的元素 即表示有环
    // node := map[*ListNode]struct{}{} // 存空结构 空间占用更小
    // for head != nil {
	// 	if _, ok := node[head]; ok { // 从map中取元素
	// 		return head // 如能取出 就返回该节点
	// 	}
	// 	node[head] = struct{}{} // 把该节点存入map
	// 	head = head.Next        // 节点后移
	// }
	// return nil

    // 快慢指针
	fast, slow := head, head              // 快，慢指针
	for fast != nil && fast.Next != nil { // 快指针走的快 故用于条件判断
		fast = fast.Next.Next // 快指针 一次走2步
		slow = slow.Next      // 慢指针 一次走1步
		if fast == slow {     // 如快慢指针相遇（即表示有环）
			for slow != head { // 慢指针不是指向的head 如是直接返回head
				// 让slow和head向后移动 当相遇时 就是环的入口
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}
