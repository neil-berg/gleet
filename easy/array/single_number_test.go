package easy

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{arr: []int{1, 2, 2}, want: 1},
		{arr: []int{3, 1, 2, 2, 1}, want: 3},
		{arr: []int{3, 3, 1, 1, 4}, want: 4},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := SingleNumber(tt.arr)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
