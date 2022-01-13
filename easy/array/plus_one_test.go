package easy

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{arr: []int{9}, want: []int{1, 0}},
		{arr: []int{8}, want: []int{9}},
		{arr: []int{8, 9}, want: []int{9, 0}},
		{arr: []int{9, 9}, want: []int{1, 0, 0}},
		{arr: []int{1, 2, 3}, want: []int{1, 2, 4}},
		{arr: []int{9, 8, 9}, want: []int{9, 9, 0}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := PlusOne(tt.arr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
