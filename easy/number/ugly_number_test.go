package easy

import "testing"

func TestUglyNumber(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{n: 2, want: true}, {n: 3, want: true}, {n: 4, want: true}, {n: 5, want: true},
		{n: 8, want: true}, {n: 9, want: true}, {n: 10, want: true}, {n: 11, want: false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := UglyNumber(tt.n)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
