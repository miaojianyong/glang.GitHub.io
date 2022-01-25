# 贪心算法

给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：
选择某个下标 i 并将 nums[i] 替换为 -nums[i]
重复这个过程恰好 k 次。可以多次选择同一个下标 i
以这种方式修改数组后，返回数组 可能的最大和

示例1：
输入：nums = [4,2,3], k = 1
输出：5
解释：选择下标 1 ，nums 变为 [4,-2,3] 。
即修改k次数组中的元素 如上 2改为-2 然后使的该数组的和是最大值
示例2：
输入：nums = [3,-1,0,2], k = 3
输出：6
解释：选择下标 (1, 2, 2) ，nums 变为 [3,1,0,2] 。
即修改数组元素 -1改为1 0改为0[即0的负数还是0] 0改为0 修改k次后 数组元素和最大

package main

import (
	"fmt"
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	// 从大到小排序数组 以绝对值方式
	// 即元素数组是[3 -1 0 2]	--> 排序后[3 2 -1 0]
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	// 遍历排序后的数组
	// 把数组中的 负数元素 改为正数
	for i := 0; i < len(nums); i++ {
		// 如k次数大于0 且 当前值小于0
		if k > 0 && nums[i] < 0 {
			nums[i] = -nums[i] // 让当前值变为其 正数
			k--                // k次数-1
		}
	}
	// 上述循环完成后，k的值=k-[数组中负数元素的个数]
	// 如 k=4,上述数组的负数元素个数为1个 故 4-1=3
	// 如 数组中负数元素的个数 大于等于k，那么循环结束后k就为0

	// 如此时的k为奇数 去修改数组最后一个元素即可[即偶数修改后还是原来的数]
	// 此时有两种情况：
	// 1. 上述数组中没有本来就没有 负数元素
	// 2. 上述数组中有 负数元素 但负数元素没有k大
	// 即 那么转变数组中数值最小的元素，即 修改数组最后一个元素即可[即偶数修改后还是原来的数]
	if k%2 == 1 {
		// 让数组的最后一个元素 执行取反操作
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	result := 0
	for i := 0; i < len(nums); i++ { // 遍历转换后的数组 使其各元素相加
		result += nums[i]
	}
	return result
}
func main() {
	fmt.Println(largestSumAfterKNegations([]int{3, -1, 0, 2}, 4))
}
