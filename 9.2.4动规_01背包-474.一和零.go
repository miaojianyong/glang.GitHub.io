## 动规_01背包

给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 

输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
输出：4
解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 
即，m个0和n个1 可组成的最大子集

package main

import "fmt"

// 动规 二维数组
func findMaxForm(strs []string, m int, n int) int {
	// 定义数组
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	// 遍历
	for i := 0; i < len(strs); i++ {
		zeroNum, oneNum := 0, 0
		//计算0,1 个数
		//或者直接strings.Count(strs[i],"0")
		for _, v := range strs[i] {
			if v == '0' {
				zeroNum++
			}
		}
		oneNum = len(strs[i]) - zeroNum
		// 从后往前 遍历背包容量
		for j := m; j >= zeroNum; j-- {
			for k := n; k >= oneNum; k-- {
				// 推导公式
				dp[j][k] = maxNum(dp[j][k], dp[j-zeroNum][k-oneNum]+1)
			}
		}
		//fmt.Println(dp)
	}
	return dp[m][n]
}

// 动规 三维数组 效率低内存消耗多
func findMaxForm2(strs []string, m int, n int) int {
	//dp的第一个index代表项目的多少，第二个代表的是背包的容量
	//所以本处项目的多少是len（strs），容量为m和n
	dp := make([][][]int, len(strs)+1)
	for i := 0; i <= len(strs); i++ {
		//初始化背包容量
		strDp := make([][]int, m+1)
		for j := 0; j < m+1; j++ {
			tmp := make([]int, n+1)
			strDp[j] = tmp
		}
		dp[i] = strDp
	}
	for k, value := range strs {
		//统计每个字符串0和1的个数
		var zero, one int
		for _, v := range value {
			if v == '0' {
				zero++
			} else {
				one++
			}
		}
		k += 1
		//计算dp
		for i := 0; i <= m; i++ {
			for j := 0; j <= n; j++ {
				//如果装不下
				dp[k][i][j] = dp[k-1][i][j]
				//如果装的下
				if i >= zero && j >= one {
					dp[k][i][j] = maxNum(dp[k-1][i][j], dp[k-1][i-zero][j-one]+1)
				}
			}
		}
	}
	return dp[len(strs)][m][n]
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	fmt.Println(findMaxForm(strs, 5, 3))  // 4
	fmt.Println(findMaxForm2(strs, 5, 3)) // 4
}
