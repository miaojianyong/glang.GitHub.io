## 贪心_两个维度权衡

n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
你需要按照以下要求，给这些孩子分发糖果：
1. 每个孩子至少分配到 1 个糖果。
2. 相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。

输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
即数组中每个元素代表一个孩子 每个孩子至少1个糖 且相邻的孩子 元素值大的比小的多1个糖

func candy(ratings []int) int {
	need := make([]int, len(ratings))
	sum := 0
	// 初始化 即每人至少1颗糖果
	for i := 0; i < len(ratings); i++ {
		need[i] = 1
	}
	// 从两个维度遍历
	// 即正序循环 和 倒序循环
	// 1. 先从左到右 当前右边大于左边 右边+1
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i+1] > ratings[i] {
			need[i+1] = need[i] + 1
		}
	}
	// 2. 再从右到左 当左边大于右边 左边+1
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			need[i-1] = findMax(need[i-1], need[i]+1)
		}
	}
	// 计算总糖果数
	for i := 0; i < len(ratings); i++ {
		sum += need[i]
	}
	return sum
}
func findMax(n1, n2 int) int { // 最大值
	if n1 > n2 {
		return n1
	}
	return n2
}
