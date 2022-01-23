## 二叉树的遍历方式

func levelOrder(root *TreeNode) [][]int {
  var res [][]int
	// 判断树是否为空
	if root == nil {
		return res
	}
	queue := list.New()   // 创建一个链表
	queue.PushBack(root)  // 把根节点存储到链表
	var arr []int         // 存储每一层的值
	for queue.Len() > 0 { // 列表的长度大于0
		queueLen := queue.Len() // 获取当前层的长度
		for i := 0; i < queueLen; i++ {
			// 把上述存储的根节点 让其出队列
			node := queue.Remove(queue.Front()).(*TreeNode)
			// 先把左节点的值存放到列表
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			// 再把右节点的值存放到列表
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			// 把该层的值存放到切片
			arr = append(arr, node.Val)
		}
		// 一层结束后 把切片结果 追加到res中
		res = append(res, arr)
		// 然后清空当前层的切片
		arr = []int{}
	}
	return res
}
