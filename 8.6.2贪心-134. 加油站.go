## 贪心

在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1

输入: 
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]
输出: 3
解释:
从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
因此，3 可为起始索引。

// 暴力解法
func canCompleteCircuit(gas, cost []int) int {
	// 遍历消耗的油量
	for i := 0; i < len(cost); i++ {
		res := gas[i] - cost[i] // 记录剩余油量
		index := (i + 1) % len(cost)
		for res > 0 && index != i { // 模拟以i为起点 汽车行驶一圈的过程
			res += gas[index] - cost[index] // 即从 大于0的出发
			index = (index + 1) % len(cost) // 更新index
		}
		// 如以i为起点 汽车行驶一圈 剩余油量 大于等于0 则返回其实位置
		if res >= 0 && index == i {
			return i
		}
	}
	return -1
}

// 贪心算法
func canCompleteCircuit2(gas, cost []int) int {
	cur := 0   // 当前消耗
	total := 0 // 总消耗
	start := 0
	for i := 0; i < len(gas); i++ {
		cur += gas[i] - cost[i]
		total += gas[i] - cost[i]
		if cur < 0 {
			// 更新起点位置 cur重新计算
			start = i + 1
			cur = 0
		}
	}
	if total < 0 { // 说明走不了一圈
		return -1
	}
	return start
}
