## 回溯_分割
一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案

输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
即把字符串分割成每个元素都是回文的集合

func partition(s string) [][]string {
	var tmpString []string // 切割字符串集合
	var res [][]string     // 结果集
	backTracking(s, tmpString, 0, &res)
	return res
}
func backTracking(s string, tmpString []string, startIndex int, res *[][]string) {
	// 如果到字符串末尾 把符合的子串 添加到结果集
	if startIndex == len(s) {
		t := make([]string, len(tmpString))
		copy(t, tmpString)
		*res = append(*res, t)
	}
	// 遍历字符串
	for i := startIndex; i < len(s); i++ {
		// 判断是否是回文子串 即在s中的 [startIndex,i]字串
		if isPartition(s, startIndex, i) { // 是回文 就添加到结果集
			tmpString = append(tmpString, s[startIndex:i+1])
		} else { // 不是就跳过
			continue
		}
		backTracking(s, tmpString, i+1, res)     // 递归
		tmpString = tmpString[:len(tmpString)-1] // 回溯
	}
}
func isPartition(s string, startIndex, end int) bool { // 判断是否是回文
	left := startIndex
	right := end
	for left < right { // 左侧索引 小于 右侧索引
		if s[left] != s[right] {
			return false
		}
		left++  // 左侧 向右移动
		right-- // 右侧 向左移动
	}
	return true
}

func main() {
	fmt.Println(partition("aab")) // [[a a b] [aa b]]
}
