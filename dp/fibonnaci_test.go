package dp

import (
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 5, want: 5},
		{n: 6, want: 8},
		{n: 50, want: 12586269025},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			cache := make(map[int]int)
			got := Fib(tt.n, cache)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
