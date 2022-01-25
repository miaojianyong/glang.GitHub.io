## 回溯_其他

给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
示例:
输入: [4, 6, 7, 7]
输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
即输出数组的所有递增子序列 且 每个子序列长度至少为2

func findSubsequences(nums []int) [][]int {
    var res [][]int
	var path []int
	var tackTrack func(int)
	tackTrack = func(start int) {
		// 当start索引到 数组nums末尾就退出
		if start > len(nums) {
			return
		}
		// 递归终止条件
		if len(path) > 1 { // 即至少是2个元素
			//temp := make([]int, len(path))
			//copy(temp, path)
			//res = append(res, temp)
			// 简写
			res = append(res, append([]int{}, path...))
		}
		// 每一层都新建一个哈希表 用于去重
		used := make(map[int]bool)
		// 遍历nums数组
		for i := start; i < len(nums); i++ {
			// 当前元素nums[i] 小于 path中的最大数 就跳过 因无法成为升序
			// 当前元素nums[i] 已经用过 也跳过
			if (len(path) > 0 && nums[i] < path[len(path)-1]) || used[nums[i]] {
				continue
			}
			// 每使用一个元素 就设置为true
			used[nums[i]] = true
			// 处理每个元素
			path = append(path, nums[i])
			tackTrack(i + 1)          // 递归
			path = path[:len(path)-1] // 回溯
		}
	}
	tackTrack(0)
	return res
}

方法二：递归枚举 + 减枝
如果有两个相同的元素，我们会考虑这样四种情况：
1. 前者被选择，后者被选择
2. 前者被选择，后者不被选择
3. 前者不被选择，后者被选择
4. 前者不被选择，后者不被选择
其中第二种情况和第三种情况其实是等价的，我们这样限制之后，舍弃了第二种，保留了第三种，于是达到了去重的目的
var (
    temp []int
    ans [][]int
)

func findSubsequences(nums []int) [][]int {
    ans = [][]int{}
    dfs(0, math.MinInt32, nums)
    return ans
}

func dfs(cur, last int, nums []int) {
    if cur == len(nums) {
        if len(temp) >= 2 {
            t := make([]int, len(temp))
            copy(t, temp)
            ans = append(ans, t)
        }
        return
    }
    if nums[cur] >= last {
        temp = append(temp, nums[cur])
        dfs(cur + 1, nums[cur], nums)
        temp = temp[:len(temp)-1]
    }
    if nums[cur] != last {
        dfs(cur + 1, last, nums)
    }
}
