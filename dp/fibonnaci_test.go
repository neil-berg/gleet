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
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Fib(tt.n)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
