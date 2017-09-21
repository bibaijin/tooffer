package q6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRebuildBinaryTree(t *testing.T) {
	cases := []struct {
		preorderList []int
		inorderList  []int
	}{
		{
			[]int{1, 2, 4, 7, 3, 5, 6, 8},
			[]int{4, 7, 2, 1, 5, 3, 8, 6},
		},
	}

	for _, c := range cases {
		b, err := RebuildBinaryTree(c.preorderList, c.inorderList)
		require.Nil(t, err)
		gotPreorderList := b.PreorderTraversal()
		require.Equal(t, len(c.preorderList), len(gotPreorderList))
		for i, v := range c.preorderList {
			require.Equal(t, v, gotPreorderList[i])
		}
		gotInorderList := b.InorderTraversal()
		require.Equal(t, len(c.inorderList), len(gotInorderList))
		for i, v := range c.inorderList {
			require.Equal(t, v, gotInorderList[i])
		}
	}
}
