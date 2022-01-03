package easy

import "testing"

func TestSameTree(t *testing.T) {
	//  a1:          b1:
	//   1             1
	// 2   3         3   2
	a1 := &Tree{Root: &TreeNode{Value: 1}}
	a1.Root.Left = &TreeNode{Value: 2}
	a1.Root.Right = &TreeNode{Value: 3}

	b1 := &Tree{Root: &TreeNode{Value: 1}}
	b1.Root.Right = &TreeNode{Value: 2}
	b1.Root.Left = &TreeNode{Value: 3}

	//  a2:           b2:
	//   1             1
	// 2                 2
	a2 := &Tree{Root: &TreeNode{Value: 1}}
	a2.Root.Left = &TreeNode{Value: 2}

	b2 := &Tree{Root: &TreeNode{Value: 1}}
	b2.Root.Right = &TreeNode{Value: 2}

	//  a2:           b2:
	//   1             1
	// 2   3         2   3
	a3 := &Tree{Root: &TreeNode{Value: 1}}
	a3.Root.Left = &TreeNode{Value: 2}
	a3.Root.Right = &TreeNode{Value: 3}

	b3 := &Tree{Root: &TreeNode{Value: 1}}
	b3.Root.Left = &TreeNode{Value: 2}
	b3.Root.Right = &TreeNode{Value: 3}

	cases := []struct {
		a, b *TreeNode
		want bool
	}{
		{a: a1.Root, b: b1.Root, want: false},
		{a: a2.Root, b: b2.Root, want: false},
		{a: a3.Root, b: b3.Root, want: true},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			got := SameTree(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
