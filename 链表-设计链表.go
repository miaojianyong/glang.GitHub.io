## 链表
描述：编写函数，如下功能
1. get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
2. addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
3. addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
4. addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
5. deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

type MyLinked struct {
	Val  int
	Next *MyLinked
}

type MyLinkedList struct {
	//链表长度
	Size int
	//头节点
	Head *MyLinked
	//尾节点
	Tail *MyLinked
}

func Constructor() MyLinkedList {
	return MyLinkedList{0, nil, nil}
}

func (this *MyLinkedList) Get(index int) int {
	//不满足条件直接返回
	if index < 0 || index >= this.Size {
		return -1
	}
	//如果index为0，直接返回头节点的值
	if index == 0 {
		return this.Head.Val
	}
	//遍历找到index节点的值
	node := this.Head
	for node.Next != nil {
		//index肯定不为0
		index--
		if node.Next != nil {
			node = node.Next
		}
		//满足条件，返回值退出
		if index == 0 {
			return node.Val
		}
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	//在头节点加入一个节点，那么这个节点就是以后的头节点了
	this.Head = &MyLinked{
		Val:  val,
		Next: this.Head,
	}
	//如果当前链表为空，那么这个增加的节点即是头节点也是尾节点
	if this.Size == 0 {
		this.Tail = this.Head
	}
	//长度+1
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	//如果当前链表为空，那么增加尾部，也就是增加头部
	if this.Size == 0 {
		this.Tail = &MyLinked{
			Val:  val,
			Next: nil,
		}
		this.Head = this.Tail
		this.Size++
		return
	}
	//增加尾节点
	this.Tail.Next = &MyLinked{
		Val:  val,
		Next: nil,
	}
	//更新尾节点
	this.Tail = this.Tail.Next
	//更新长度
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		//不插入
		return
	}
	if index == this.Size { //添加到末尾
		this.AddAtTail(val)
		return
	}
	if index <= 0 {
		//头部插入节点
		this.AddAtHead(val)
	}
	node := this.Head
	for node.Next != nil {
		index--
		if index == 0 {
			newNode := &MyLinked{
				Val:  val,
				Next: node.Next,
			}
			node.Next = newNode
			this.Size++
			return
		}
		node = node.Next
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	//如果等于0，就是删除头节点，链表长度-1
	if index == 0 {
		this.Head = this.Head.Next
		this.Size--
		return
	}
	node := this.Head
	for node.Next != nil {
		index--
		if index == 0 {
			if node.Next.Next == nil {
				//到了最后一个节点，删除最后一个节点
				node.Next = nil
				this.Tail = node
				this.Size--
				return
			}
			//其他情况就是删除中间一个节点
			node.Next = node.Next.Next
			this.Size--
			return
		}
		node = node.Next
	}
}
