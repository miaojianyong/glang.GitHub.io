## 字符串
输入：s = ["h","e","l","l","o"]
输出：["o","l","l","e","h"]
即反转数组中个元素位置

func reverseString(s []byte)  {
    // 双指针
    left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left] // 交换
		left++                                // 右移
		right--                               // 左移
	}
}
