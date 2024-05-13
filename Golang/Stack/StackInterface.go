package Stack

type Stack interface {
	Push(v any)
	Pop() any
	Peek() any
	IsEmpty() bool
	Size() int
	Clear()
	Print()
}
