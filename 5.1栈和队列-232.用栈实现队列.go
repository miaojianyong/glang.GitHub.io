## 栈和队列
使用栈实现队列的下列操作：
push(x) -- 将一个元素放入队列的尾部。
pop() -- 从队列首部移除元素。
peek() -- 返回队列首部的元素。
empty() -- 返回队列是否为空。

package main

type MyQueue struct {
	inStack  []int // 入栈 即用于压入元素
	outStack []int // 出栈 即用于输出元素
}

func Constructor() MyQueue { // 构造函数
	return MyQueue{ // 返回 初始化各字段
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) { // 把元素添加到队尾
	// 1. 先把 出栈 里面的各个元素 从后到前 挨个追加到 入栈 中
	for len(this.outStack) != 0 { // 如 出栈 不为空
		val := this.outStack[len(this.outStack)-1]           // 获取 出栈 的最后一个元素
		this.outStack = this.outStack[:len(this.outStack)-1] // 出栈 长度-1
		this.inStack = append(this.inStack, val)             // 把上述在 出栈 中获取的元素 追加到 入栈 中
	}
	// 上述操作后 出栈 已为空
	// 2. 再把 元素x 追加到 入栈 中，即实现 在队尾添加元素
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int { // 弹出队头的元素
	// 1. 先把 入栈 里面的各个元素 从后到前 挨个追加到 出栈 中
	for len(this.inStack) != 0 { // 如 入栈 不为空
		val := this.inStack[len(this.inStack)-1]          // 获取 入栈 的最后一个元素
		this.inStack = this.inStack[:len(this.inStack)-1] // 入栈 长度-1
		this.outStack = append(this.outStack, val)        // 把上述在 入栈 中获取的元素 追加到 出栈 中
	}
	// 上述操作后 入栈 已为空
	if len(this.outStack) == 0 {
		return 0
	}
	// 2. 再从 出栈 中获取其最后的元素(该元素就是队头的元素，即入栈最里面的元素)
	val := this.outStack[len(this.outStack)-1]
	// 3. 把 出栈 长度-1
	this.outStack = this.outStack[:len(this.outStack)-1]
	return val // 返回该元素
}

func (this *MyQueue) Peek() int { // 获取对头元素
	// 1. 先把 入栈 里面的各个元素 从后到前 挨个追加到 出栈 中
	for len(this.inStack) != 0 { // 如 入栈 不为空
		val := this.inStack[len(this.inStack)-1]          // 获取 入栈 的最后一个元素
		this.inStack = this.inStack[:len(this.inStack)-1] // 入栈 长度-1
		this.outStack = append(this.outStack, val)        // 把上述在 入栈 中获取的元素 追加到 出栈 中
	}
	// 上述操作后 入栈 已为空
	if len(this.outStack) == 0 {
		return 0
	}
	// 2. 再从 出栈 中获取其最后的元素(该元素就是队头的元素，即入栈最里面的元素)
	val := this.outStack[len(this.outStack)-1]
	// 该函数 仅仅是查看队头的元素 故不用 删除上述元素
	return val // 返回该元素
}

func (this *MyQueue) Empty() bool { // 队列是否为空
	// 即看 入栈 和 出栈 两个栈是否为空
	return len(this.inStack) == 0 && len(this.outStack) == 0
}
