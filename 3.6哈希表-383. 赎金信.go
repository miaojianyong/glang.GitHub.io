## 哈希表
输入：ransomNote = "a", magazine = "b"
输出：false
即字符串magazine中的字符 是否能组成字符串ransomNote

func canConstruct(ransomNote string, magazine string) bool {
	// 判断 字符串magazine里面 是否存在 字符串ransomNote
	//return strings.Contains(magazine, ransomNote)
	// 不能用这个内置函数 因为还要去顺序 如aab和baa
	// 上述会函数会返回false 因本题不要求顺序 只是看后者字符串是否能组成前者 故应返回true

	if len(ransomNote) > len(magazine) {
		return false
	}
	cnt := [26]int{}
	for _, ch := range magazine {
		cnt[ch-'a']++ // 对应字符 +a
	}
	for _, ch := range ransomNote {
		cnt[ch-'a']--        // 对应字符 -a 即如有相同字符应为0
		if cnt[ch-'a'] < 0 { // 如小于0 即该字符串有上述字符串不同的字符
			return false
		}
	}
	return true
}
