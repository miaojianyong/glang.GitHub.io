## 栈和队列
输入: ["2", "1", "+", "3", " * "]
输出: 9
解释: 该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
即把数组的元素转换成算术表达式

package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{} // 定义栈
	for _, token := range tokens {
		// 把当前字符串转数字
		val, err := strconv.Atoi(token)
		if err == nil { // 如没错 即表示当前字符串可转数字
			stack = append(stack, val) // 把当前数字追加到栈
		} else { // 否则就是运算符
			// 从栈中取出两个数 倒是第2个数字给n1 最后1个给n2
			// 即先入栈的数 作为左操作数 后入栈的作为右操作数
			n1, n2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2] // 更新栈的长度
			switch token {               // 匹配运算符 并把结果入栈
			case "+":
				stack = append(stack, n1+n2)
			case "-":
				stack = append(stack, n1-n2)
			case "*":
				stack = append(stack, n1*n2)
			case "/":
				stack = append(stack, n1/n2)
			}
		}
	}
	return stack[0] // 返回栈的第1个元素
}
