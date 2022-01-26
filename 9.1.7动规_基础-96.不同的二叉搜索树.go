## 动规_基础

一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数

输入：n = 3
输出：5
即，从1到n，可组成几种不同的二叉搜索数组 如下所示：
1      1        2         3    3
 \      \      / \       /    /
  3      2    1   3     2    1
 /        \            /      \
2          3          1        2  -->即 1-3可组成这 5种不同的二叉搜索树
即公式为：dp[3] = dp[2]*dp[0] + dp[1]*dp[1] + dp[0]*dp[2]
故递推公式：dp[i] += dp[j-1] * dp[i-j]
j-1 是以j为头节点的 左子树 节点的数量
i-j 是以j为头节点的 右子树 节点的数量

// 动规
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
