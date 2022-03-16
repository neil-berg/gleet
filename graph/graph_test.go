package graph

import (
	"testing"
)

func TestHasPath(t *testing.T) {
	graph := &Graph{
		AdjList: AdjacencyList{
			1: []int{2, 4},
			2: []int{3},
			3: []int{},
			4: []int{2, 5, 6},
			5: []int{},
			6: []int{},
		},
	}

	tests := []struct {
		src, dest int
		want      bool
	}{
		{src: 1, dest: 5, want: true},
		{src: 2, dest: 5, want: false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			gotDFS := graph.HasPathDFS(tt.src, tt.dest)
			gotBFS := graph.HasPathBFS(tt.src, tt.dest)
			if gotDFS != tt.want || gotBFS != tt.want {
				t.Errorf("got %v, want %v", gotDFS, tt.want)
			}
		})
	}
}
