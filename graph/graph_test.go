package graph

import (
	"fmt"
	"testing"
)

func newGraph() *Graph {
	g := &Graph{}
	g.Init()
	return g
}

func TestCreateAdjacency(t *testing.T) {
	g := newGraph()
	edges := [][]int{
		[]int{1, 2},
		[]int{1, 3},
		[]int{3, 4},
	}
	g.CreateAdjacency(edges)
	fmt.Println(g.Adjacency)
}

func TestTraverse(t *testing.T) {
	g := newGraph()
	g.Adjacency = map[int][]int{
		1: []int{2, 3},
		2: []int{1, 3},
		3: []int{1, 4, 5},
		4: []int{4},
		5: []int{3, 6},
		6: []int{5, 7},
		7: []int{6},
	}

	// values := g.TraverseDepthFirstRecur(1)
	// values2 := g.TraverseDepthFirst(1)
	values := g.TraverseBreadth(1)
	fmt.Println(values)
	// fmt.Println(values)
	// fmt.Println(values2)
}

func TestHasPath(t *testing.T) {
	g := newGraph()
	g.Adjacency = map[int][]int{
		1: []int{2, 3},
		2: []int{1, 3},
		3: []int{1, 4, 5},
		4: []int{4},
		5: []int{3},
		6: []int{7},
		7: []int{6},
	}

	tests := []struct {
		src, dst int
		want     bool
	}{
		{src: 1, dst: 5, want: true},
		{src: 4, dst: 7, want: false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := g.HasPathBFS(tt.src, tt.dst)
			// got := g.HasPathDFS(tt.src, tt.dst)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnectedComponents(t *testing.T) {
	g := newGraph()
	edges := [][]int{
		[]int{1, 2},
		[]int{3, 4},
		[]int{3, 5},
		[]int{3, 6},
		[]int{3, 7},
		[]int{8, 9},
	}
	g.CreateAdjacency(edges)
	got := g.ConnectedComponents()
	want := 3
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
