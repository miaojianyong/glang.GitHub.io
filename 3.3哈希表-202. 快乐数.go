## 哈希表
描述：判断一个数 n 是不是快乐数
输入：n = 19
输出：true
解释：
1² + 9² = 82
8² + 2² = 68
6² + 8² = 100
1² + 0² + 0² = 1
即一个数各个位进行平方然后相加，其结果各个位再进行平方相加，最后等于1

package main

import "fmt"

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] { // 如当前数不为1 且 该数没有在map中
		// 让该数去进行 个位平方后相加操作
		// 如返回的结果在map中存在就返回该数
		n, m[n] = getSum(n), true
	}
	fmt.Println(n, m)
	return n == 1
}

func isHappy2(n int) bool {
	// 快慢指针 一个一次走1步 一个一次走两步
	// 一个指针指向当前值 一个指针指向其各位平方后的和 的结果
	slow, fast := n, getSum(n) // 2 4
	// 如该该值各位平方后的和 的结果 不为1 且 两个指针执行的值不同
	for fast != 1 && slow != fast {
		fmt.Println(slow, fast)
		// 两个指针各进行运算 如两个指针指向相同值就退出
		slow = getSum(slow)         // 4 走1步
		fast = getSum(getSum(fast)) // 37 // 走2步
	}
	return fast == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10) // 从个位数开始计算平方
		n = n / 10                 // 下一位，如十位
	}
	return sum
}
func main() {
	fmt.Println(isHappy(2)) // false
	// n=4 map[2:true 4:true 16:true 20:true 37:true 42:true 58:true 89:true 145:true]
	// 即145 的结果是 1²+4²+5²=42 在map已存在 故返回 145 != 1 返回false
	fmt.Println(isHappy(7)) // true
	// n=1 map[7:true 10:true 49:true 97:true 130:true]
	// 即10 的结果是 1²+0²=1 因 1 == 1 故返回true
}
