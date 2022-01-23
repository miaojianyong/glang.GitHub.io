## 栈和队列
使用队列实现栈的下列操作：
push(x) -- 元素 x 入栈
pop() -- 移除栈顶元素
top() -- 获取栈顶元素
empty() -- 返回栈是否为空

package main

// 1. 用两个队列
//type MyStack struct {
//	// queue1 第1个元素表示栈顶
//	queue1 []int // 用于所有数据
//	queue2 []int // 用于入栈时存储数据
//}
//
//func Constructor1() MyStack {
//	return MyStack{ // 初始化
//		queue1: make([]int, 0),
//		queue2: make([]int, 0),
//	}
//}
//
//func (this *MyStack) Push(x int) { // 入栈
//	// 1. 先把元素追加到 queue2 里
//	this.queue2 = append(this.queue2, x)
//	// 2. 如queue1中有元素 把元素从前到后 依次追加到queue2中
//	// 即实现 新加入的元素x在queue2的最前面 即栈顶位置
//	for len(this.queue1) != 0 { // 如 queue1 不为空
//		// 把queue1中所有元素 从前面取出[即前面的是栈顶 故从栈顶开始取] 依次 追加到queue2中
//		this.queue2 = append(this.queue2, this.queue1[0]) // 取出queue1第1个元素
//		this.queue1 = this.queue1[1:]                     // 更新queue1的长度
//	}
//	// 上述完成 queue1就成空的了
//	// 3. 让queue1和queue2交换 即queue2空了 其里面的元素 交换 到了queue1中
//	// 即实现 新加入的元素x在queue1的最前面 即栈顶位置
//	this.queue1, this.queue2 = this.queue2, this.queue1
//}
//
//func (this *MyStack) Pop() int { // 出栈
//	val := this.queue1[0]         // 即queue1的第1个元素 是要出栈的元素
//	this.queue1 = this.queue1[1:] // 去除第1个元素 即更新queue1的长度
//	return val
//}
//
//func (this *MyStack) Top() int { // 获取栈顶元素
//	return this.queue1[0] // 直接返回queue1的第1个元素即可
//}
//
//func (this *MyStack) Empty() bool { // 栈是否为空
//	return len(this.queue1) == 0
//}

// 2. 用1个队列实现
type MyStack struct {
	queue []int
}

func Constructor1() MyStack {
	return MyStack{ // 初始化
		queue: make([]int, 0),
	}
}
func (this *MyStack) Push(x int) { // 入栈
	n := len(this.queue)               // 先获取原队列长度
	this.queue = append(this.queue, x) // 把元素追加到 队列末尾
	// 然后把队列前面的元素取出 依次追加到 队列末尾 实现上述x在队列头部 即栈顶位置
	for n != 0 {
		this.queue = append(this.queue, this.queue[0]) // 取第1个元素
		this.queue = this.queue[1:]                    // 更新队列长度
		n--                                            // 即仅把原队列里面的元素取出依次追加到队列尾部即可
	}
}
func (this *MyStack) Pop() int { // 出栈
	val := this.queue[0]        // 队列第1个元素 即栈顶元素
	this.queue = this.queue[1:] // 更新队列长度
	return val
}
func (this *MyStack) Top() int { // 获取栈顶元素
	return this.queue[0] // 直接返回第1个元素即可
}

func (this *MyStack) Empty() bool { // 栈是否为空
	return len(this.queue) == 0
}
func main() {
	s1 := Constructor1()
	s1.Push(2)
}
