## 字符串
输入: s = "abcdefg", k = 2
输出: "bacdfeg"
即从0开始 每隔k个位置 反转对应字符串

package main

import "fmt"

func reverseStr(s string, k int) string {
	// 双指针法反转字符串函数
	var reverse func(s []byte)
	reverse = func(s []byte) {
		l, r := 0, len(s)-1
		for l < r {
			s[l], s[r] = s[r], s[l]
			l++
			r--
		}
	}
	str := []byte(s) // 把字符串转byte切片
	n := len(s)
	for i := 0; i < n; i += 2 * k { // 每次+2k
		// 如到k的位置 小于等于 n 就返回k前面的字符
		if i+k <= n {
			reverse(str[i : i+k])
		} else { // 否则 就是小于n 那么后面的字符全反转
			reverse(str[i:n])
		}
	}
	return string(str) // 转成字符串返回
}

func main() {
	fmt.Println(reverseStr("asdfigjkl", 3))  // dsafiglkj
	fmt.Println(reverseStr("asdfigjk", 3))   // dsafigkj
	fmt.Println(reverseStr("asdfigjklw", 3)) // dsafiglkjw
}
