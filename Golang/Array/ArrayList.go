package Array

// List 定义一个通用的列表结构
type List interface {
	Add(item interface{})
	Get(index int) interface{}
	Size() int
}

// ArrayList 结构
type ArrayList struct {
	items []interface{}
}

// NewArrayList 创建一个新的 ArrayList
func NewArrayList() *ArrayList {
	return &ArrayList{items: make([]interface{}, 0)}
}

// Add 添加元素到 ArrayList
func (list *ArrayList) Add(item interface{}) {
	list.items = append(list.items, item)
}

// Get 获取指定索引的元素
func (list *ArrayList) Get(index int) interface{} {
	if index < 0 || index >= len(list.items) {
		return nil
	}
	return list.items[index]
}

// Size 返回列表大小
func (list *ArrayList) Size() int {
	return len(list.items)
}
