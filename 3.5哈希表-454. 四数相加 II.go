## 哈希表
输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
输出：2
解释：
两个元组如下：
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0
即4个数组有几个是满足个元素相加等于0的

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    m := make(map[int]int,len(nums1)*len(nums2)) // 数组长度都相等 故map的最大长度就是数组的平方
	count := 0
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			// 把nums1、nums2的元素和作为key,该和出现的次数作为value存储到map
			m[v1+v2]++
		}
	}
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			// 统计nums3、nums4元素和的 负数为key，在上面map中的次数
			count += m[-(v3 + v4)] // m[-v3-v4]
		}
	}
	return count
}
