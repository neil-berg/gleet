package medium

import (
	"reflect"
	"testing"
)

func TestAtoi(t *testing.T) {
	tests := map[string]struct {
		s    string
		want int32
	}{
		"leading whitespace": {s: "   100", want: int32(100)},
		"negative":           {s: "-100", want: int32(-100)},
		"positive":           {s: "100", want: int32(100)},
		"bad prefix":         {s: "abcd-100", want: int32(0)},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			got := Atoi(test.s)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
