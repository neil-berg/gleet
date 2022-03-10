package easy

import (
	"reflect"
	"testing"
)

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{arr: []int{1, 2, 2, 1, 1, 3}, want: true},
		{arr: []int{1, 2}, want: false},
		{arr: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, want: true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := UniqueOccurrences(tt.arr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
