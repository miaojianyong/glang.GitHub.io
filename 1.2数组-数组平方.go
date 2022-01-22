## 数组
描述：把数组的每个元素平方后，新数组为升序数组
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
func sortedSquares(nums []int) []int {
  n := len(nums)
	i, j, k := 0, n-1, n-1
	res := make([]int, n)
	for i <= j {
    // 左边、右边元素平方
		left, rigth := nums[i]*nums[i], nums[j]*nums[j]
    // 然后比较 使得新数组为升序
		if left > rigth {
			res[k] = left
			i++
		} else {
			res[k] = rigth
			j--
		}
		k--
	}
	return res
}
