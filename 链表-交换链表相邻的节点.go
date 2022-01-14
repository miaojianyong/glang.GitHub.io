## 链表
描述：交换链表其中相邻的节点
输入：head = [1,2,3,4]
输出：[2,1,4,3]

func swapPairs(head *ListNode24) *ListNode24 {
	// 递归
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//next := head.Next                // 保存下个节点
	//head.Next = swapPairs(next.Next) // 递归交换下个节点
	//next.Next = head                 // 把当前节点给下个节点
	//return next

	// 迭代 效率更高
	dummy := &ListNode24{} // 虚拟头节点
	dummy.Next = head      // 让虚拟节点指向头节点
	pre := dummy           // 定义指针 指向虚拟节点
	// 如头节点不为空 且它的下个节点也不为空 就进行下述交换操作
	for head != nil && head.Next != nil {
		next := head.Next     // 定义变量 保存头节点的下个节点 如2
		head.Next = next.Next // 头节点1 开始指向2 改为 指向3
		next.Next = head      // next=2 原指向3 改为 指向 1
		pre.Next = next       // pre原指向1 改为指向 2
		// 即原来是 pre->1->2->3->4 上述操作后 pre->2->1->3->4

		// 更新指针位置
		pre = head       // 即 2->1[pre]->3->4
		head = head.Next // 即 2->1[pre]->3[head]->4
	}
	return dummy.Next
}
