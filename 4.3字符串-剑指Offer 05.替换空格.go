## 字符串
示例 1： 输入：s = "We are happy."
输出："We%20are%20happy."
即把字符串里面的空格 替换成对应的字符

package main

func replaceSpace(s string) string {
	// 把字符串 s 中的每个空格替换成"%20"
	// 如 "We are happy." -> "We%20are%20happy."
	// 方式1 遍历添加 即遍历字符串把 空格 替换成%20
	/*b := []byte(s)
	result := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' { // 如发现空白字符串把 %20 添加到结果集
			result = append(result, []byte("%20")...)
		} else { // 否则把当前字符添加到结果集
			result = append(result, b[i])
		}
	}
	return string(result)*/

	// 方式2 原地修改
	b := []byte(s)
	length := len(b)
	spaceCount := 0
	// 计算空格数量
	for _, v := range b {
		if v == ' ' {
			spaceCount++
		}
	}
	// 扩展原有切片
	resizeCount := spaceCount * 2
	tmp := make([]byte, resizeCount)
	b = append(b, tmp...)
	i := length - 1
	j := len(b) - 1
	for i >= 0 {
		if b[i] != ' ' {
			b[j] = b[i]
			i--
			j--
		} else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			i--
			j = j - 3
		}
	}
	return string(b)
}
