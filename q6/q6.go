package q6

import (
	"fmt"

	"github.com/bibaijin/tooffer/tree"
)

// RebuildBinaryTree 重建二叉树
func RebuildBinaryTree(preorderList []int, inorderList []int) (*tree.BinaryTree, error) {
	if len(preorderList) == 0 {
		return nil, nil
	}

	if len(preorderList) == 1 {
		return &tree.BinaryTree{
			Value: preorderList[0],
			Left:  nil,
			Right: nil,
		}, nil
	}

	index, err := search(inorderList, preorderList[0])
	if err != nil {
		return nil, err
	}

	left, err := RebuildBinaryTree(preorderList[1:index+1], inorderList[0:index])
	if err != nil {
		return nil, err
	}

	right, err := RebuildBinaryTree(preorderList[index+1:], inorderList[index+1:])
	if err != nil {
		return nil, err
	}

	return &tree.BinaryTree{
		Value: preorderList[0],
		Left:  left,
		Right: right,
	}, nil
}

func search(xs []int, a int) (int, error) {
	for i, x := range xs {
		if x == a {
			return i, nil
		}
	}

	return 0, fmt.Errorf("%d not found in %v", a, xs)
}
