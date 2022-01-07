package easy

import "testing"

func TestRemoveDupInPlace(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{
		{a: []int{1, 2, 2, 3}, want: 3},
		{a: []int{1, 2, 2, 3, 3, 4}, want: 4},
		{a: []int{1, 2, 2, 3, 3, 4, 4}, want: 4},
		{a: []int{1, 2, 2, 3, 3, 4, 4, 5}, want: 5},
		{a: []int{1, 1, 1, 1, 1, 1}, want: 1},
		{a: []int{}, want: 0},
		{a: []int{1}, want: 1},
		{a: []int{1, 2}, want: 2},
		{a: []int{1, 2, 3}, want: 3},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := RemoveDupInPlace(tt.a)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
