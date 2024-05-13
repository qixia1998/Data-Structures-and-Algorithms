package Stack

import "fmt"

/*
 基于链表实现的Stack
*/

type node struct {
	item any
	next *node
}

type LinkedListStack struct {
	topPtr *node
	count  int
}

func (list *LinkedListStack) Push(item any) {
	list.topPtr = &node{item: item, next: list.topPtr}
	list.count++
}

func (list *LinkedListStack) Size() int {
	return list.count
}

func (list *LinkedListStack) IsEmpty() bool {
	return list.count == 0
}

func (list *LinkedListStack) Clear() {
	list.topPtr = nil
	list.count = 0
}

func (list *LinkedListStack) Pop() any {
	if list.IsEmpty() {
		return nil
	}
	result := list.topPtr.item
	list.topPtr = list.topPtr.next
	list.count--
	return result
}

func (list *LinkedListStack) Peek() any {
	if list.IsEmpty() {
		return nil
	}
	return list.topPtr.item
}

func (list *LinkedListStack) Print() {
	if list.IsEmpty() {
		fmt.Println("stack is empty")
	}
	for current := list.topPtr; current != nil; current = current.next {
		fmt.Println(current.item)
	}
}
