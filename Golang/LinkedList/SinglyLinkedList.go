package LinkedList

import "fmt"

// 单链表
type Ordered interface {
	~string | ~int | ~float64
}

type Node[T Ordered] struct {
	Item T
	next *Node[T]
}

type LinkedList[T Ordered] struct {
	first       *Node[T]
	numberItems int
}

// 遍历
func (l *LinkedList[T]) traverse(head *Node[T]) {
	p := head
	for p != nil {
		fmt.Println(p.Item)
		p = p.next
	}
}

// 查找
func (l *LinkedList[T]) find(value T) *Node[T] {
	head := new(Node[T])
	p := head // p用来遍历链表结点，初始化为head
	for p != nil {
		if p.Item == value {
			return p
		}
		p = p.next
	}
	return nil
}

// 插入 头部插入、尾部插入、给定结点之后插入

// 链表头部插入
func (l *LinkedList[T]) insertAtHead(value T) {
	newNode := &Node[T]{
		value,
		nil,
	}
	head := new(Node[T])
	newNode.next = head
	head = newNode
}

// 链表尾部插入 顺序
func (l *LinkedList[T]) insertAtTail(value T) {
	newNode := &Node[T]{
		value,
		nil,
	}
	head := new(Node[T])
	if head == nil {
		head = newNode
	} else {
		p := head
		for p != nil {
			p = p.next
		}
		p.next = newNode
	}
}

// 链表尾部插入 优化
func (l *LinkedList[T]) insertAtTail2(value T) {
	head := new(Node[T])
	tail := new(Node[T]) // 记录链表尾部结点位置
	newNode := &Node[T]{
		value,
		nil,
	}
	if head == nil {
		head = newNode
		tail = newNode
	} else {
		tail.next = newNode
		tail = newNode
	}
}

// 链表尾部插入 引入虚拟结点
func (l *LinkedList[T]) insertAtTail3(value T) {
	head := new(Node[T])
	tail := head

	newNode := &Node[T]{
		value,
		nil,
	}
	tail.next = newNode
	tail = newNode
}

// 给定结点之后插入
func (l *LinkedList[T]) insertAfter(p *Node[T], value T) {
	if p == nil {
		return
	}
	newNode := &Node[T]{
		value,
		nil,
	}
	newNode.next = p.next
	p.next = newNode
}

// 删除给定结点之后的节点
func (l *LinkedList[T]) deleteNextNode(p *Node[T]) {
	if p == nil || p.next == nil {
		return
	}
	p.next = p.next.next
}
