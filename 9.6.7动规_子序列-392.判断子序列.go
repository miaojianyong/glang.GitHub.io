## 动规_子序列 编辑距离

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）

输入：s = "abc", t = "ahbgdc"
输出：true
即看s是否是t的子序

// 双指针法
func isSubsequence(s string, t string) bool {
	x,y:=len(s),len(t)
    i,j:=0,0 // 双指针
    for i<x && j<y{
        // 如 s的当前字符 等于 t的当前字符
        if s[i] == t[j] {
            i++ // 去比较s的下个字符
        }
        j++ // 否则 去匹配t的下个字符
    }
    return i == x
}

// 用内置还是判断其索引
func isSubsequence(s string, t string) bool {
	for i := 0; i < len(s); i++ {
		// 如果 s的当前字符在t中 就返回其在t中的索引 否则返回-1
		j := strings.Index(t, string(s[i]))
		if j == -1 {
			return false
		}
		t = t[j+1:] // 切割t,即删除t当前索引前面的所有元素(包含当前元素)
		// 然后继续循环切割后的字符串
	}
	return true
}

// 动规 -- 相比上述两个方法内存消耗更少 执行效率都差不多
func isSubsequence(s string, t string) bool {
    dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			// 如果 s中和t中的相同索引对应的值 相同
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1 // 长度+1
			} else { // 否则 去t的下个元素匹配
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(s)][len(t)] == len(s)
}
