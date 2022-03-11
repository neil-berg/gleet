package dp

import "testing"

func TestCanSum(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   bool
	}{
		{arr: []int{2, 3}, target: 7, want: true},
		{arr: []int{5, 3, 4, 7}, target: 7, want: true},
		{arr: []int{2, 4}, target: 7, want: false},
		{arr: []int{2, 3, 5}, target: 8, want: true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			cache := make(map[int]bool)
			got := CanSum(tt.arr, tt.target, cache)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
