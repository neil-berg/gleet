package easy

import "testing"

func TestValidParens(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{s: "()[]{}", want: true},
		{s: "((({{}})))", want: true},
		{s: "({)", want: false},
		{s: "(]", want: false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := ValidParens(tt.s)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
