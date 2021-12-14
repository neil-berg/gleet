package easy

import "testing"

func TestMinSumByDigits(t *testing.T) {
	tests := map[string]struct {
		arr  []int
		want int
	}{
		"test 1": {arr: []int{6, 8, 4, 5, 2, 3}, want: 604},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			got := MinSumByDigits(test.arr)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
