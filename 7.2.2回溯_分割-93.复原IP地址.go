## 回溯_分割
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔

输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]
输入：s = "0000"
输出：["0.0.0.0"]
即把字符串分割成有效的IP地址

func restoreIpAddress(s string) []string {
	var res, path []string
	backTracking(s, path, 0, &res)
	return res
}
func backTracking(s string, path []string, startIndex int, res *[]string) {
	// 终止条件
	if startIndex == len(s) && len(path) == 4 {
		/*tmpIpString := path[0] + "." + path[1] + "." + path[2] + "." + path[3] // 拼接路径
		*res = append(*res, tmpIpString)                                       // 追加到结果集*/

		tmp := strings.Join(path, ".")
		*res = append(*res, tmp)

		//*res = append(*res, strings.Join(path, "."))
	}
	// 遍历字符串
	for i := startIndex; i < len(s); i++ {
		path = append(path, s[startIndex:i+1]) // 分隔字符串 追加到path
		if i-startIndex+1 <= 3 && len(path) <= 4 && isNormalIp(s, startIndex, i) {
			backTracking(s, path, i+1, res) // 递归
		} else { // 如首尾[path字串]超过3个，或路径对于4个，或不是合法地址 就退出
			return
		}
		path = path[:len(path)-1] // 回溯
	}
}

func isNormalIp(s string, start, end int) bool { // 判断地址是否合法
	val, _ := strconv.Atoi(s[start : end+1]) // 转成int
	// 如果s长度大于1 且第1位是0 就不合法
	if end-start+1 > 1 && s[start] == '0' {
		return false
	}
	if val > 255 {
		return false
	}
	return true
}

// 方式2 -- 效率更高
func restoreIpAddresses(s string) []string {
	res := []string{}
	var dfs func(subRes []string, start int)

	dfs = func(subRes []string, start int) {
		if len(subRes) == 4 && start == len(s) {
			res = append(res, strings.Join(subRes, "."))
			return
		}
		if len(subRes) == 4 && start < len(s) {
			return
		}
		for length := 1; length <= 3; length++ {
			if start+length-1 >= len(s) {
				return
			}
			if length != 1 && s[start] == '0' {
				return
			}
			str := s[start : start+length]
			if n, _ := strconv.Atoi(str); n > 255 {
				return
			}
			subRes = append(subRes, str)
			dfs(subRes, start+length)
			subRes = subRes[:len(subRes)-1]
		}
	}
	dfs([]string{}, 0)
	return res
}
