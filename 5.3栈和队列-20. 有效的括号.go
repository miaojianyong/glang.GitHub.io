## 栈和队列
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
即判断字符串中的括号是否是正确的

package main

func isValid(s string) bool {
	if len(s)%2 == 1 { // 如字符串长度不是偶数 返回false
		return false
	}
	// 定义字符集
	hash := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0)
	for i := range s {
		// 如果是下述字符的一种 就把该字符 追加到栈
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			// 如 栈不是空的 且 栈的最后一个元素 等于 字符串中的元素 即匹配到一个括号
		} else if len(stack) > 0 && stack[len(stack)-1] == hash[s[i]] {
			// 匹配到 就删除栈最后一个元素 即栈顶元素
			stack = stack[:len(stack)-1] // 长度-1 即去除最后一个元素
		} else {
			return false
		}
	}
	return len(stack) == 0
}
