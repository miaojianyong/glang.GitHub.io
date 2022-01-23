## 字符串
示例 1: 输入: haystack = "hello", needle = "ll" 输出: 2
示例 2: 输入: haystack = "aaaaa", needle = "bba" 输出: -1
即输出字串在字符串中的其实索引位置

package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// 方法一:前缀表使用减1实现

// getNext 构造前缀表next
// params:
//		  next 前缀表数组
//		  s 模式串
func getNext1(next []int, s string) { // 获取子串的前后缀数组
	j := -1 // j表示 最长相等前后缀长度
	next[0] = j

	for i := 1; i < len(s); i++ {
		for j >= 0 && s[i] != s[j+1] { // 前后缀不同
			j = next[j] // 回退前一位
		}
		if s[i] == s[j+1] { // 前后缀相同
			j++
		}
		// j(前缀长度)赋值给next[i] 要记录相同前后缀的长度
		next[i] = j // next[i]是i（包括i）之前的最长相等前后缀长度
	}
}
func strStr1(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := make([]int, len(needle))
	getNext1(next, needle)
	j := -1 // 模式串的起始位置 next为-1 因此也为-1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j] // 寻找下一个匹配点
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == len(needle)-1 { // j指向了模式串的末尾
			return i - len(needle) + 1
		}
	}
	return -1
}

// 方法二: 前缀表无减一或者右移

// getNext 构造前缀表next
// params:
//		  next 前缀表数组
//		  s 模式串
func getNext2(next []int, s string) {
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}
func strStr2(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	j := 0
	next := make([]int, n)
	getNext2(next, needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] // 回退到j的前一位
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))  // 2
	fmt.Println(strStr("hello", "lle")) // -1
	fmt.Println(strStr("", ""))         // 0

	fmt.Println(strStr1("hello", "ll")) // 2
	fmt.Println(strStr1("hello", "lo")) // 3
}
