package LinkedList

import "fmt"

// 单链表
type LinkedList struct {
	head *Node
	count int
}

type Node struct {
	value any
	next *Node
}

func (list *LinkedList) Size() int {
	return list.count
}

func (list *LinkedList) IsEmpty() bool {
	return (list.count == 0)
}

func (list *LinkedList) AddHead(value any) {
	list.head = &Node{value, list.head}
	list.count++
}

func (list *LinkedList) AddTail(value any) {
	curr := list.head
	newNode := &Node{value, nil}

	if curr == nil{
		list.head = newNode
		return
	}

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

// 指定位置插入
func (list *LinkedList) InsertAt(index int, value any) error {
	if index < 1 || index > list.count + 1{
		return fmt.Errorf("Index out of bounds error")
	}
	newNode := Node{value, nil}

	if index == 0 {
		newNode.next = list.head
		list.head = &newNode
		list.count++
		return nil
	}
	node := list.head
	num := 0
	previous := node
	for num < index {
		previous = node
		num++
		node = node.next
	}
	newNode.next = node
	previous.next = &newNode
	list.count++
	return nil
}

func (list *LinkedList) Print() {
	temp := list.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println("")
}