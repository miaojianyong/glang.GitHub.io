## 贪心_区间问题

输入：points = [[10,16],[2,8],[1,6],[7,12]]
输出：2
解释：对于该样例，x = 6 可以射爆 [2,8],[1,6] 两个气球，以及 x = 11 射爆另外两个气球
即用最少的箭可以接触到数组中 各个元素区间的值

func findMinArrowShots(points [][]int) int {
    res := 1 // 箭数
	// 按照二维数组的第1个 排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	for i := 1; i < len(points); i++ {
		// 如 前1位的右边 小于 后1位的左边 就不重叠
		if points[i-1][1] < points[i][0] {
			res++
		} else { // 否则就是重叠了
			// 更新重叠气球的 最右边 也就是下1个的 最右边的值
			points[i][1] = func(a, b int) int {
				if a > b {
					return b
				}
				return a
				// 即在 前1个的最右边 和 后1个的最右边 取最小值
			}(points[i-1][1], points[i][1])
		}
	}
	return res
}
