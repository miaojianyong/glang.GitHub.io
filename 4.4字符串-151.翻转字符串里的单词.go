## 字符串
输入: "the sky is blue"
输出: "blue is sky the"
即反转字符串各个单词 以及去除单词间多余的空格

package main

import "fmt"

func reverseWords(s string) string {
	// 方式1
	// 反转字符串函数
	var reverse func(b *[]byte, left, right int)
	reverse = func(b *[]byte, left, right int) {
		for left < right {
			(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
			left++
			right--
		}
	}
	//1.使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	//删除头部冗余空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	//删除单词间冗余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	//删除尾部冗余空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	//2.反转整个字符串
	reverse(&b, 0, len(b)-1)
	//3.反转单个单词  i单词开始位置，j单词结束位置
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(&b, i, j-1)
		i = j
		i++
	}
	return string(b)
}

func reverseWords2(s string) string {
	// 三部曲：1.翻转字符串，2.找单词左右边界并翻转单词，3.拼接字符串
	b := []byte(s)
	var reverse func(start, end int)
	reverse = func(start, end int) {
		for start < end {
			b[start], b[end] = b[end], b[start]
			start++
			end--
		}
	}
	// 1.翻转整个字符串
	reverse(0, len(s)-1)
	// 2.确定单词边界并翻转单个单词
	start, end, i := 0, 0, 0
	// 拼接结果变量
	ans := ""
	// 遍历s的byte数组
	for ; i < len(b); i++ {
		// 每轮循环处理一个单词
		// 通过遇到非空字符作为单词的开始边界
		if b[i] != ' ' {
			start = i
			for i < len(b) && b[i] != ' ' { // 得到完整单词
				i++
			}
			// 单词的右边界
			end = i - 1
			// 左右边界已经确定，翻转单个单词
			reverse(start, end)
			// 3.拼接到结果string中, 因为每个单词需用空格分隔 故需在前面加空格拼接
			ans += " " + string(b[start:end+1])
		}
	}
	// 去掉头部空格
	return ans[1:]
}

func main() {
	fmt.Println(reverseWords(" hello  world! "))
	fmt.Println(reverseWords2(" hello  world! "))
}
