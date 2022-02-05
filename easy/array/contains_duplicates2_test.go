package easy

import "testing"

func TestContainsDuplicates2(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		want bool
	}{
		{arr: []int{1, 2, 3, 1}, k: 3, want: true},
		{arr: []int{1, 0, 1, 1}, k: 1, want: true},
		{arr: []int{1, 2, 3, 1, 2, 3}, k: 2, want: false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := ContainsDuplicates2(tt.arr, tt.k)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
