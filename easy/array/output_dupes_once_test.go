package easy

import (
	"reflect"
	"testing"
)

func TestDuplicatesOnce(t *testing.T) {
	tests := []struct {
		arr, want []int
	}{
		{arr: []int{1, 1, 2, 2, 2}, want: []int{1, 2}},
		{arr: []int{1, 1, 1, 2, 2, 2, 3}, want: []int{1, 2}},
		{arr: []int{1, 1, 1, 2, 2, 2, 3, 3}, want: []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := DuplicatesOnce(tt.arr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
