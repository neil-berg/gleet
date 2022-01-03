package easy

// SortedArrayToBST converts a sorted int array to a height-balanced binary search tree.
func SortedArrayToBST(arr []int) *TreeNode {
	var recur func(start, end int) *TreeNode

	recur = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := (start + end) / 2
		node := &TreeNode{Value: arr[mid]}
		node.Left = recur(0, mid-1)
		node.Right = recur(mid+1, end)
		return node
	}

	root := recur(0, len(arr)-1)
	return root
}
