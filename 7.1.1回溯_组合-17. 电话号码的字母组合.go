## 回溯_组合
一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
即上述数组元素对应的字符集是电话 9宫格按键上对应的字母
然后输入对应的数字，让其该数字上对应的字母返回其组合

func combine(digits string) []string {
	lenth := len(digits)
	if lenth == 0 || lenth > 4 {
		return nil
	}
	digitsMap := [10]string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}
	res := make([]string, 0)
	backTrack("", digits, 0, digitsMap, &res)
	return res
}

// 临时字符集 参数字符串[0-9] 每个字符对应的字符集 结果集
func backTrack(tempString, digits string, index int, digitsMap [10]string, res *[]string) {
	if len(tempString) == len(digits) { // 终止条件 拼接的字符串长度 等于 digits长度
		*res = append(*res, tempString)
		return
	}
	tmpK := digits[index] - '0'        // 把参数转int 如 '2'-> 2
	letter := digitsMap[tmpK]          // 取出参数对应的字符集 如 2 -> "abc"
	for i := 0; i < len(letter); i++ { // 遍历字符对应的字符集 如遍历"abc"
		tempString = tempString + string(letter[i])            // 拼接结果
		backTrack(tempString, digits, index+1, digitsMap, res) // 递归
		tempString = tempString[:len(tempString)-1]            // 回溯
	}
}
