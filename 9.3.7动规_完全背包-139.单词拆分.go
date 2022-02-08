## 动规_完全背包

给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成
即：数组中的元素是否能拼接成字符串s

func wordBreak(s string, wordDict []string) bool {
	dictMap := make(map[string]bool)
	for _, v := range wordDict {
		dictMap[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			tempWord := s[j:i]
			if dp[j] && dictMap[tempWord] { // 是否在单词表
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
