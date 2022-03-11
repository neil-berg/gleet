package dp

import "testing"

func TestGridTraveller(t *testing.T) {
	tests := []struct {
		m, n, want int
	}{
		{m: 1, n: 1, want: 1},
		{m: 2, n: 2, want: 2},
		{m: 3, n: 2, want: 3},
		{m: 3, n: 3, want: 6},
		{m: 18, n: 18, want: 2333606220},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			cache := make(map[string]int)
			got := GridTraveller(tt.m, tt.n, cache)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
