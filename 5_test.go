package tooffer

import (
	"testing"

	"github.com/bibaijin/tooffer/list"
)

func TestPrintListReversely(t *testing.T) {
	node3 := list.SingleListNode{
		Value: 3,
		Next:  nil,
	}
	node2 := list.SingleListNode{
		Value: 2,
		Next:  &node3,
	}
	node1 := list.SingleListNode{
		Value: 1,
		Next:  &node2,
	}
	node0 := list.SingleListNode{
		Value: 0,
		Next:  &node1,
	}
	cases := []struct {
		in *list.SingleListNode
	}{
		{&node0},
	}

	for _, c := range cases {
		printListReversely(c.in)
		printListReverselyRecursively(c.in)
	}
}
