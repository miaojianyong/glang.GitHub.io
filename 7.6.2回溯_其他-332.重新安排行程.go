## 回溯_其他

一份航线列表 tickets ，其中 tickets[i] = [fromi, toi] 表示飞机出发和降落的机场地点。请你对该行程进行重新规划排序
输入：tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
输出：["JFK","MUC","LHR","SFO","SJC"]
即如下图：
MUC --> LHR --> SFO
 ^               |
 |               V
JFK             SJC
即按照上述数组中的集合 排列出 飞机飞行的轨迹图

思路：
1. 从起点出发，进行深度优先搜索。
2. 每次沿着某条边从某个顶点移动到另外一个顶点的时候，都需要删除这条边。
3. 如果没有可移动的路径，则将所在节点加入到栈中，并返回。
var d map[string][]string
var ans []string

func findItinerary(tickets [][]string) []string {
	d = map[string][]string{}
	for _, v := range tickets {
		d[v[0]] = append(d[v[0]], v[1])
	}
	for _, v := range d {
		sort.Strings(v)
	}
	ans = []string{}
	dfs("JFK")
	return ans
}

func dfs(f string) {
	for len(d[f]) > 0 {
		v := d[f][0]
		d[f] = d[f][1:]
		dfs(v)
	}
	ans = append([]string{f}, ans...)
}
