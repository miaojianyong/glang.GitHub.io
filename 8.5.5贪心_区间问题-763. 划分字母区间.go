## 贪心_区间问题

字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表

输入：S = "ababcbacadefegdehijhklij"
输出：[9,7,8]
解释：
划分结果为 "ababcbaca", "defegde", "hijhklij"。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少

func partitionLabels(s string) []int {
    var res []int;
    var marks [26]int;
    size, left, right := len(s), 0, 0;
    for i := 0; i < size; i++ {
        marks[s[i] - 'a'] = i;
    }
    for i := 0; i < size; i++ {
        right = max(right, marks[s[i] - 'a']);
        if i == right {
            res = append(res, right - left + 1);
            left = i + 1;
        }
    }
    return res;
}

func max(a, b int) int {
    if a < b {
        a = b;
    }
    return a;
}
