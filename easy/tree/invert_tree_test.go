package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvertTree(t *testing.T) {
	tree := &Tree{Root: &TreeNode{Value: 1}}
	tree.Root.Left = &TreeNode{Value: 2}
	tree.Root.Right = &TreeNode{Value: 3}
	tree.Root.Left.Left = &TreeNode{Value: 4}
	tree.Root.Left.Right = &TreeNode{Value: 5}
	tree.Root.Right.Left = &TreeNode{Value: 6}
	tree.Root.Right.Right = &TreeNode{Value: 7}

	root := InvertTree(tree.Root)
	assert.Equal(t, root.Value, 1)
	assert.Equal(t, root.Left.Value, 3)
	assert.Equal(t, root.Right.Value, 2)
	assert.Equal(t, root.Left.Left.Value, 7)
	assert.Equal(t, root.Left.Right.Value, 6)
	assert.Equal(t, root.Right.Left.Value, 5)
	assert.Equal(t, root.Right.Right.Value, 4)
}
