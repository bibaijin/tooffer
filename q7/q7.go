package q7

import (
	"github.com/bibaijin/tooffer/stack"
)

// Queue 表示队列
type Queue struct {
	a stack.Stack
	b stack.Stack
}

// Push 插入元素
func (q *Queue) Push(x interface{}) {
	q.a.Push(x)
}

// Pop 取出元素
func (q *Queue) Pop() interface{} {
	if q.b.IsEmpty() {
		for !q.a.IsEmpty() {
			x := q.a.Pop()
			q.b.Push(x)
		}
	}

	return q.b.Pop()
}

// IsEmpty 判读是否为空
func (q Queue) IsEmpty() bool {
	return q.b.IsEmpty() && q.a.IsEmpty()
}
