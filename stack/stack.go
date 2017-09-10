package stack

// Stack 表示栈
type Stack interface {
	Push(v interface{})
	Pop() interface{}
	IsEmpty() bool
}

type sliceStack struct {
	slice []interface{}
}

// New 返回初始化后的 Stack
func New() Stack {
	slice := make([]interface{}, 0)
	return &sliceStack{
		slice: slice,
	}
}

func (s *sliceStack) Push(v interface{}) {
	s.slice = append(s.slice, v)
}

func (s *sliceStack) Pop() (x interface{}) {
	x, s.slice = s.slice[len(s.slice)-1], s.slice[:len(s.slice)-1]
	return x
}

func (s sliceStack) IsEmpty() bool {
	return len(s.slice) == 0
}
