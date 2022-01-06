package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedArrayToBST(t *testing.T) {
	root := SortedArrayToBST([]int{-10, -3, 0, 5, 9})
	assert.Equal(t, root.Value, 0)
	assert.Equal(t, root.Left.Value, -10)
	assert.Equal(t, root.Left.Right.Value, -3)
	assert.Equal(t, root.Right.Value, 5)
	assert.Equal(t, root.Right.Right.Value, 9)
}
