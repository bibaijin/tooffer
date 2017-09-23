package queue

// Queue 表示队列
type Queue interface {
	Push(x interface{})
	Pop() interface{}
	IsEmpty() bool
}

// sliceQueue 表示 slice 实现的队列
type sliceQueue struct {
	xs []interface{}
}

// New 返回初始化后的 Queue
func New() Queue {
	xs := make([]interface{}, 0)
	return &sliceQueue{
		xs: xs,
	}
}

// Push 插入元素
func (q *sliceQueue) Push(x interface{}) {
	q.xs = append(q.xs, x)
}

// Pop 取出元素
func (q *sliceQueue) Pop() (x interface{}) {
	if q.IsEmpty() {
		return nil
	}

	x, q.xs = q.xs[0], q.xs[1:]
	return x
}

func (q sliceQueue) IsEmpty() bool {
	return len(q.xs) == 0
}
