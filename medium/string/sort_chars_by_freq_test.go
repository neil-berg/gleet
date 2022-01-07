package medium

import "testing"

func TestSortCharsByFreq(t *testing.T) {
	tests := []struct{ s, want string }{
		{s: "tree", want: "eetr"},
		{s: "trrree", want: "rrreet"},
		{s: "treetree", want: "eeeettrr"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := SortCharsByFreq(tt.s)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
