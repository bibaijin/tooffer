package tooffer

import (
	"fmt"

	"github.com/bibaijin/tooffer/list"
	"github.com/bibaijin/tooffer/stack"
)

func printListReversely(l *list.SingleListNode) {
	s := stack.New()
	for n := l; n != nil; n = n.Next {
		s.Push(n.Value)
	}
	for !s.IsEmpty() {
		x := s.Pop()
		fmt.Printf("%v\n", x)
	}
}

func printListReverselyRecursively(l *list.SingleListNode) {
	if l != nil {
		printListReverselyRecursively(l.Next)
		fmt.Printf("%v\n", l.Value)
	}
}
