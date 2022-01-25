## 贪心

给定一个二叉树，我们在树的节点上安装摄像头。
节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。
计算监控树的所有节点所需的最小摄像头数量

输入：[0,0,null,0,0]
输出：1
解释：如图所示，一台摄像头足以监控所有节点。
如图：
      O
     /
    X  --> 即在此处安装摄像头 可监听所有结点 即返回1 就是1个摄像头
   / \
  O   O

const inf = math.MaxInt64 / 2 // 常量
func minCameraCover(root *TreeNode) int {
    var dfs func(*TreeNode) (a, b, c int)
    dfs = func(node *TreeNode) (a, b, c int) {
        if node == nil {
            return inf, 0, 0
        }
        lefta, leftb, leftc := dfs(node.Left)
        righta, rightb, rightc := dfs(node.Right)
        a = leftc + rightc + 1
        b = min(a, min(lefta+rightb, righta+leftb))
        c = min(a, leftb+rightb)
        return
    }
    _, ans, _ := dfs(root)
    return ans
}

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}
