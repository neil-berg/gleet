package easy

import (
	"reflect"
	"testing"
)

func TestReverseInt(t *testing.T) {
	cases := map[string]struct {
		n    int
		want int
	}{
		"positive":     {n: 271, want: 172},
		"negative":     {n: -123, want: -321},
		"leading zero": {n: 90, want: 9},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got := ReverseInt(c.n)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got: %v, want: %v", got, c.want)
			}
		})
	}
}
