package Array

import (
	"errors"
	"fmt"
)

type Array struct {
	data   []any
	length uint
}

// NewArray 初始化数组
func (a *Array) NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]any, capacity),
		length: 0,
	}
}

func (a *Array) Len() uint {
	return a.length
}

// IsIndexOutOfRange 判断数组索引是否越界
func (a *Array) IsIndexOutOfRange(index uint) bool {
	if index >= uint(cap(a.data)) {
		return true
	}
	return false
}

// Find 通过索引查找数组，索引范围[0,n-1]
func (a *Array) Find(index uint) (any, error) {
	if a.IsIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return a.data[index], nil
}

// Insert 在索引index上插入值
func (a *Array) Insert(index uint, v any) error {
	if a.Len() == uint(cap(a.data)) {
		return errors.New("array is full")
	}
	if index != a.length && a.IsIndexOutOfRange(index) {
		return errors.New("out of index range")
	}

	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = v
	a.length++
	return nil
}

// Delete 删除索引index上的值
func (a *Array) Delete(index uint) (any, error) {
	if a.IsIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	v := a.data[index]
	for i := index; i < a.Len()-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return v, nil
}

// Print 打印数组
func (a *Array) Print() {
	var format string
	for i := uint(0); i < a.Len(); i++ {
		format += fmt.Sprintf("|%+v", a.data[i])
	}
	fmt.Println(format)

}
