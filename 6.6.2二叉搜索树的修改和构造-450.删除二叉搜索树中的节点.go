## 二叉搜索树的修改和构造
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。
返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：
首先找到需要删除的节点；
如果找到了，删除它。

输入：root = [5,3,6,2,4,null,7], key = 3
输出：[5,4,6,2,null,null,7]
解释：给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/delete-node-in-a-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
如：
      5
   /     \
  3       6
 / \       \
2   4       7 --> 删除节点3
输出：
      5
   /     \
  4       6
 /         \
2           7
或：
      5
   /     \
  2       6
   \       \
    4       7

// 递归
func deleteNode(root *TreeNode, key int) *TreeNode {
	// 如果是空数 返回空
	if root == nil {
		return nil
	}
	// 如果找到要删除的节点 中
	if root.Val == key {
		// 当前节点没有左子树 就让右子树代替root
		if root.Left == nil {
			return root.Right
		}
		// 当前节点没有右子树 就让左子树代替root
		if root.Right == nil {
			return root.Left
		}
		// 如果当前节点 左右子树都有 // 有两个子节点，寻找左边或者右边接替自己，这里直接进行数据的交换就行
		// 找后继节点
		// 即找到 右子树的 最左节点
		next := root.Right     // 给当前节点的右节点
		for next.Left != nil { // 左节点不为空 就一直向下找
			next = next.Left
		}
		// 然后递归删除旧数据
		root.Val = next.Val
		root.Right = deleteNode(root.Right, next.Val)
		return root

		/*// 否则说明左右子树都不为空,去找左子树中的最大值对应的节点
		leftMax := root.Left       // 先移动到左子树
		for leftMax.Right != nil { // 如果左子树的右子树不为空
			leftMax = leftMax.Right // 就一直移动到它的右子树直到为空
		} // 找到左子树的最大值节点后，把当前根节点的右子树作为它的右子树
		leftMax.Right = root.Right
		root = root.Left // 然后把根节点移动到左子树*/

		/*// 找右子树的最小值
		rightMin := root.Right
		for rightMin.Left != nil {
			rightMin = rightMin.Left
		}
		rightMin.Left = root.Left
		root = root.Right*/
	}
	// 如当前节点 大于 key就去左子树
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else { // 否则就去 右子树
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

// 方法2
func deleteNode2(root *TreeNode, key int) *TreeNode {
	if root == nil { // 如果当前节点为空
		return nil // 就返回空指针
	}
	if key < root.Val { // 如果要删除的值小于根节点的值
		// 就递归去左子树上删除节点，并且返回的二叉树是当前根节点的左子树
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val { // 如果要删除的值大于根节点的值
		// 就递归去右子树上删除节点，并且返回的二叉树是当前根节点的右子树
		root.Right = deleteNode(root.Right, key)
	} else { // 如果不是以上两种情况，说明要删除的值等于当前根节点的值
		if root.Left == nil { // 先判断当前根节点是否只有一个子树
			return root.Right // 是的话就返回那个唯一子树即可
		} else if root.Right == nil {
			return root.Left
		}
		rightMin := root.Right
		for rightMin.Left != nil {
			rightMin = rightMin.Left
		}
		rightMin.Left = root.Left
		root = root.Right
	}
	return root // 最后返回当前根节点即可
}
