package easy

import "testing"

func TestSearchInsertPosition(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   int
	}{
		{arr: []int{1, 3, 5, 6}, target: 5, want: 2},
		{arr: []int{1, 3, 5, 6}, target: 2, want: 1},
		{arr: []int{1, 3, 5, 6}, target: 7, want: 4},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := SearchInsertPosition(tt.arr, tt.target)
			if got != tt.want {
				t.Errorf("got %v, wnt %v", got, tt.want)
			}
		})
	}
}
