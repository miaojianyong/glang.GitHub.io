## 哈希表
描述：给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
示例 1: 输入: s = "anagram", t = "nagaram" 输出: true
即看两个字符串里面的字符是否相同

package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) { // 如果两个字符串长度不同 肯定不是异位词
		return false
	}
	//s1, s2 := []byte(s), []byte(t) // 把字符串转成byte切片
	//// 然后使用函数把上述切片排序 升序
	//sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	//sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	//return string(s1) == string(s2) // 转成字符串 看是否相等

	//// 哈希 数组 执行效率最高
	//strArr := [26]int{}
	//// 遍历s和t的每个字符 做+-记录
	//for _, v := range s { // 做+操作
	//	strArr[v-rune('a')]++
	//}
	//for _, v := range t { // 做-操作
	//	strArr[v-rune('a')]--
	//}
	//// 看strArr是否等于默认值
	//return strArr == [26]int{}

	// 哈希 map
	strMap := map[rune]int{}
	for _, v := range s { // 遍历s字符串的每个字符
		strMap[v]++ // 记录每个字符的个数
	}
	for _, v := range t { // 遍历t字符串的每个字符
		strMap[v]-- // 去掉每个字符的个数 即针对上述记录的相应字符
		// 如在上述map中，有元素key对应的字符个数小于0
		// 表示该字符不在s中 或 出现次数过多
		if strMap[v] < 0 {
			return false
		}
	}
	return true
}
func main() {
	s := "hello"
	t := "oellh"
	fmt.Println(isAnagram(s, t)) // true
}
