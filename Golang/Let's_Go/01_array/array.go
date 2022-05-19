package main

/**
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据类型是interface类型
 */

type Array struct {
	data   []interface{}
	length int
}

// 数组初始化内存
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]interface{}, capacity, capacity),
		length: 0,
	}
}

// 判断索引是否越界
func (array *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(array.data)) {
		return true
	}
	return false
}

// 通过索引查找数组，索引范围[0, n-1],（未查找到，返回-1）
func (array *Array) Find(value interface{}) int {
	for i := 0; i < array.length; i++ {
		if array.data[i] == value {
			return i
		}
	}
	return -1
}

// 获取数组容量
func (array *Array) GetCapacity() int {
	return cap(array.data)
}

// 获取数组长度
func (array *Array) GetLength() int {
	return array.length
}

// 判断数组是否为空
func (array *Array) IsEmpty() bool {
	return array.length == 0
}
