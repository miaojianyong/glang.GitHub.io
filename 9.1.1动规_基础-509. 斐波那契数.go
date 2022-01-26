## 动规_基础

斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
F(0) = 0，F(1) = 1
F(n) = F(n - 1) + F(n - 2)，其中 n > 1
给定 n ，请计算 F(n)

输入：n = 2
输出：1
解释：F(2) = F(1) + F(0) = 1 + 0 = 1
即当前数是前两个数的和

package main

import "fmt"

// 方式1 递归
func fib1(n int) int {
	if n < 2 {
		return n
	}
	return fib1(n-1) + fib1(n-2)
}

// 带 备忘录 递归法
func fib4(n int) int {
	if n == 0 {
		return 0
	}
	memo := make([]int, n+1) // 将备忘录 全部初始化为0
	return helper(memo, n)   // 进行带备忘录的递归
}
func helper(memo []int, n int) int {
	// base case 基础情况 即简单的问题
	if n == 1 || n == 2 {
		return 1
	}
	// 已经计算过
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = helper(memo, n-1) + helper(memo, n-2)
	return memo[n]
}

// 动态规划
func fib2(n int) int {
	if n < 2 {
		return n
	}
	// dp五步曲
	dp := make([]int, n+1) // 1. dp数组
	//2.公式 dp[i] = dp[i-1]+dp[i-2]
	//dp = [2]int{0, 2} // 3. 初始化dp数组
	// 或下述方法 初始化
	dp[0] = 0
	dp[1] = 1
	// 遍历数组 从前向后
	for i := 2; i <= n; i++ { // 因为0和1初始化了 故从2开始
		dp[i] = dp[i-1] + dp[i-2] // 套公式
	}
	return dp[n]
}

func fib3(n int) int {
	// 动态规划
	// 边界条件
	if n < 2 {
		return n
	}
	// dp五部曲
	var sum int
	// 1.定义dp数组, dp[1]表示n-1的fib值，dp[0]表示n-2的fib值
	dp := [2]int{}
	// 2.递归公式 sum = dp[0] + dp[1], dp[0] = dp[1], dp[1] = sum
	// 3.dp数组初始化, dp[0]=0, dp[1]=1
	dp = [2]int{0, 1}
	// 4.确定遍历顺序，从前向后遍历
	for i := 2; i <= n; i++ {
		// 2.递推公式
		sum = dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = sum
	}
	return dp[1]
}

// dp数组 的迭代法
func fib5(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	fmt.Println(fib1(10)) // 55
	fmt.Println(fib2(10)) // 55
	fmt.Println(fib3(10)) // 55
	fmt.Println(fib4(10)) // 55
	fmt.Println(fib5(10)) // 55
}
