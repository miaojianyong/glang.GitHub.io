## 栈和队列
输入："abbaca"
输出："ca"
即输出字符串中相邻的重复字符 如上述bb就是相邻重复项
故删除后得到"aaca" 次数aa也是相邻重复项 故也需删除 故输出"ca"

func removeDuplicates(s string) string {
    var stack []byte
    for i := 0; i < len(s);i++ {
        // 栈不空 且 当前元素与栈顶元素相等 即表示是相同字符
        if len(stack) > 0 && stack[len(stack)-1] == s[i] {
            // 弹出栈顶元素 并 忽略当前元素(s[i]) 
            // 即栈顶元素出栈 且 当前元素不入栈
            stack = stack[:len(stack)-1]
        }else{
            // 否则 当前元素入栈
            stack = append(stack, s[i])
        }
    }
    return string(stack)
}
