package tree

import "testing"

// Helper to setup a test tree:
//             5
//           /   \
//          3     7
//         / \   / \
//        2   4 6   10
func setupTree() *BST {
	bst := &BST{}
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(7)
	bst.Insert(10)
	bst.Insert(6)
	return bst
}

func TestMaxDepth(t *testing.T) {
	bst := setupTree()
	got := MaxDepth(bst.root)
	want := 3
	if got != want {
		t.Errorf("Incorrect max depth, got: %v, want %v", got, want)
	}
}

func TestMaxValue(t *testing.T) {
	bst := setupTree()
	got, _ := bst.MaxValue()
	want := 10
	if got != want {
		t.Errorf("Incorrect max value, got: %v, want %v", got, want)
	}
}

func TestMinValue(t *testing.T) {
	bst := setupTree()
	got, _ := bst.MinValue()
	want := 2
	if got != want {
		t.Errorf("Incorrect min value, got: %v, want %v", got, want)
	}
}

func TestValueExists(t *testing.T) {
	bst := setupTree()

	tests := map[string]struct {
		v    int
		want bool
	}{
		"exists":         {v: 4, want: true},
		"does not exist": {v: 1, want: false},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			got := bst.ValueExists(test.v)
			if got != test.want {
				t.Errorf("Incorrect value exists, got: %v, want %v", got, test.want)
			}
		})
	}
}
