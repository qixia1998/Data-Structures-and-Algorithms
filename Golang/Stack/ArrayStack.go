package Stack

import (
	"errors"
	"fmt"
)

/*
 基于Array 实现 Stack
*/

type ArrayStack struct {
	array []any
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		array: make([]any, 0),
	}
}

// IsEmpty 判断是否为空
func (s *ArrayStack) IsEmpty() bool {
	return len(s.array) == 0
}

func (s *ArrayStack) Push(value any) {
	s.array = append(s.array, value)
}

func (s *ArrayStack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	result := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return result, nil
}

// Peek 查看栈顶元素
func (s *ArrayStack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.array[len(s.array)-1], nil
}

func (s *ArrayStack) Clear() {
	s.array = make([]any, 0)
}

func (s *ArrayStack) Size() int {
	return len(s.array)
}

func (s *ArrayStack) Print() {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
	}
	for _, value := range s.array {
		fmt.Println(value)
	}
}
