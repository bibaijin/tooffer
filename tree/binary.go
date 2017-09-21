package tree

// BinaryTree 表示二叉树
type BinaryTree struct {
	Value interface{}
	Left  *BinaryTree
	Right *BinaryTree
}

// PreorderTraversal 先序遍历
func (b *BinaryTree) PreorderTraversal() []interface{} {
	xs := make([]interface{}, 0)

	xs = append(xs, b.Value)
	if b.Left != nil {
		xs = append(xs, b.Left.PreorderTraversal()...)
	}
	if b.Right != nil {
		xs = append(xs, b.Right.PreorderTraversal()...)
	}

	return xs
}

// InorderTraversal 中序遍历
func (b *BinaryTree) InorderTraversal() []interface{} {
	xs := make([]interface{}, 0)

	if b.Left != nil {
		xs = append(xs, b.Left.InorderTraversal()...)
	}
	xs = append(xs, b.Value)
	if b.Right != nil {
		xs = append(xs, b.Right.InorderTraversal()...)
	}

	return xs
}

// PostorderTraversal 后序遍历
func (b *BinaryTree) PostorderTraversal() []interface{} {
	xs := make([]interface{}, 0)

	if b.Left != nil {
		xs = append(xs, b.Left.PostorderTraversal()...)
	}
	if b.Right != nil {
		xs = append(xs, b.Right.PostorderTraversal()...)
	}
	xs = append(xs, b.Value)

	return xs
}
