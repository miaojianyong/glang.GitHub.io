## 字符串
示例 1:
输入: "abab"
输出: True
解释: 可由子字符串 "ab" 重复两次构成。
即判断字符串是否是重复的字串构成

package main

import (
	"fmt"
)

func repeatedSubstringPattern(s string) bool {
	// 1. 枚举
	//n := len(s)
	//for i := 1; i*2 <= n; i++ {
	//	if n%i == 0 {
	//		flag := true
	//		for j := i; j < n; j++ {
	//			if s[j] != s[j-i] {
	//				flag = false
	//				break
	//			}
	//		}
	//		if flag {
	//			return true
	//		}
	//	}
	//}
	//return false

	// 2. 新建一个新字符串str=s+s，把str的首元素和尾元素去掉，剩下的部分如果还含有s，则返回true
	//var str1 string = s + s
	//var str2 string = str1[1 : len(str1)-1]
	//fmt.Println(str1, str2)
	//if strings.Contains(str2, s) {
	//	return true
	//} else {
	//	return false
	//}
	// 简写
	// s1 := s + s
	// s2 := s1[1:len(s1) - 1]
	//return strings.Contains(s2,s)

	// 3. KMP算法
	n := len(s)
	if n == 0 {
		return false
	}
	next := make([]int, n)
	j := -1
	next[0] = j
	for i := 1; i < n; i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j
	}
	// 如 next数组最后1个元素不等于-1 且 数组长度-最长相等前后缀长度 可被数组长度整除
	// next[n-1]+1 最长相同前后缀的长度
	if next[n-1] != -1 && n%(n-(next[n-1]+1)) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(repeatedSubstringPattern("abab"))  // true
	fmt.Println(repeatedSubstringPattern("ababa")) // false
}
