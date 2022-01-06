package easy

import (
	"reflect"
	"testing"
)

func TestPalindomeNumber(t *testing.T) {
	cases := map[string]struct {
		n    int
		want bool
	}{
		"negative":       {n: -11, want: false},
		"ends with zero": {n: 10, want: false},
		"palindrome":     {n: 123321, want: true},
		"not palindrome": {n: 123322, want: false},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got := PalindromeNumber(c.n)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}
