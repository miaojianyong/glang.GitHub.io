## 贪心_区间问题

以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6]

// 贪心
func merge(intervals [][]int) [][]int {
	// 按左边界 从小到大排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals)-1; i++ { // 当前和后1个比较 用i:=0;len(xxx)-1;i++
		// 如 当前区间的 右边界 大于等于 下个区间的 左边界
		if intervals[i][1] >= intervals[i+1][0] {
			// 取 当前区间的右边界的值 和 下个去右边界的值 的最大值 给当前区间的右边边界
			intervals[i][1] = func(a, b int) int {
				if a > b {
					return a
				}
				return b
			}(intervals[i][1], intervals[i+1][1])
			// 然后把 下个区间后面的元素添加到 下个区间的位置 即i+2和后续的 添加到i+1位置
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			// 压缩长度
			i--
		}
	}
	return intervals
}

// 优化
func merge(intervals [][]int) [][]int {
    var max func(a,b int)int
    max = func(a,b int) int {
        if a > b {
            return a
        }
        return b
    }
    res := [][]int{}
    if len(intervals) <= 1 {
		return intervals
	}

	// 第一个数字升序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if right >= intervals[i][0] {
			// 有交叉
			right = max(right, intervals[i][1])
		} else {
			res = append(res, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		}
	}
	res = append(res, []int{left, right})
	return res
}
