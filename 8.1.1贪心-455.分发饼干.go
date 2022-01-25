## 贪心
假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干

输入: g = [1,2,3], s = [1,1]
输出: 1
解释: 
你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
所以你应该输出1
即把上述数组g中的孩子即胃口数 用s数组的饼干数组来分配 看有几个孩子能分配到合理的饼干

func findContentChildren(g, s []int) int {
	// 排序孩子和饼干数组
	sort.Ints(g)
	sort.Ints(s)
	child := 0 // 满足胃口的孩子数
	// child < len(g) && sIndex < len(s) child小于孩子数 sIndex小于饼干数
	for sIndex := 0; child < len(g) && sIndex < len(s); sIndex++ {
		// 如当前 饼干 大于等于 当前孩子的胃口
		if s[sIndex] >= g[child] {
			child++ // 满足胃口的孩子+1
		}
	}
	return child
}

func findContentChildren2(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	/*i, j, res := 0, 0, 0
	for i < len(g) && j < len(s) { // i小于孩子长度 j小于饼干长度
		if s[j] >= g[i] { // 当前饼干 满足 当前孩子
			res = res + 1 // 结果+1
			j++           // 下个饼干
			i++           // 下个孩子
		} else { // 否则 找下个 饼干
			j++
		}
	}
	return res*/

	// 代码简化
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] { // 当前孩子 小于等于 当前饼干
			i++
		}
		j++ // 否则 下个饼干
	}
	return i
}
