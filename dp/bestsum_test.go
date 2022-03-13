package dp

import (
	"reflect"
	"testing"
)

func reverseArr(arr []int) []int {
	dst := make([]int, len(arr))
	copy(dst, arr)

	for i := 0; i < len(dst)/2; i++ {
		j := len(dst) - i - 1
		dst[i], dst[j] = dst[j], dst[i]
	}
	return dst
}

func TestBestSum(t *testing.T) {
	tests := []struct {
		target int
		arr    []int
		want   []int
	}{
		{target: 8, arr: []int{2, 3, 5}, want: []int{3, 5}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			cache := make(map[int][]int)
			got := BestSum(tt.target, tt.arr, cache)
			ok := !reflect.DeepEqual(got, tt.want) || !reflect.DeepEqual(got, reverseArr(tt.want))
			if !ok {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
