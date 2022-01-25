贪心_序列

给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增

示例1：
输入: N = 10
输出: 9
示例2：
输入: N = 1234
输出: 1234
示例3：
输入: N = 332
输出: 299
即是整数n变为每个位的数是递增[或相等的]的

package main

import (
	"fmt"
	"strconv"
)

func monotoneIncreasingDigits(n int) int { // 运行超时
	// 检测当前数是否是递增的
	var checkNum func(int) bool
	checkNum = func(num int) bool {
		max := 10
		for num != 0 {
			t := num % 10 // 3,0,1 即数字的倒序
			if max >= t {
				max = t
			} else {
				return false
			}
			num = num / 10
		}
		return true
	}
	for i := n; i > 0; i-- {
		// 如当前数是递增的 返回当前数 否则当前数-1[即i--] 继续检测
		if checkNum(i) {
			return i
		}
	}
	return 0
}

func monotoneIncreasingDigits2(n int) int {
	if n == 0 {
		return 0
	}
	// 将数组转换成字符串
	s := strconv.Itoa(n)
	ss := []byte(s) // 把字符串转为byte数组 方便更改
	numStr := len(ss)
	if numStr <= 1 { // 个位数直接返回
		return numStr
	}
	for i := numStr - 1; i > 0; i-- { // 倒序
		// 前一个 大于 后一个 就前一个减1 后面全是9
		if ss[i-1] > ss[i] {
			ss[i-1] -= 1
			for j := i; j < numStr; j++ { // 后面全设置为9
				ss[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(ss))
	return res
}
