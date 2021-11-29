package bst

import (
	"reflect"
	"testing"
)

// Helper to setup a test tree:
//             5
//           /   \
//          3     7
//         / \   / \
//        2   4 6   10
func setupTree() *BST {
	bst := &BST{}
	for _, v := range []int{5, 3, 2, 4, 7, 10, 6} {
		bst.Insert(v)
	}
	return bst
}

func TestMaxDepth(t *testing.T) {
	bst := setupTree()
	got := MaxDepth(bst.Root)
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
				t.Errorf("Incorrect value %v exists, got: %v, want %v", test.v, got, test.want)
			}
		})
	}
}

func TestInOrderTraversal(t *testing.T) {
	bst := setupTree()
	got := bst.InOrderTraversal()
	want := []int{2, 3, 4, 5, 6, 7, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPreOrderTraversal(t *testing.T) {
	bst := setupTree()
	got := bst.PreOrderTraversal()
	want := []int{5, 3, 2, 4, 7, 6, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPostOrderTraversal(t *testing.T) {
	bst := setupTree()
	got := bst.PostOrderTraversal()
	want := []int{2, 4, 3, 6, 10, 7, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
