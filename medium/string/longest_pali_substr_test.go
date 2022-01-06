package medium

import "testing"

func TestLongestPaliSubstr(t *testing.T) {
	tests := map[string]struct {
		s    string
		want string
	}{
		"first":  {s: "abbaz", want: "abba"},
		"second": {s: "zabbaz", want: "zabbaz"},
		"third":  {s: "zabbazz", want: "zabbaz"},
		"fourth": {s: "zzabbazz", want: "zzabbazz"},
		"fifth":  {s: "zzzz", want: "zzzz"},
		"sixth":  {s: "zazbaaz", want: "zaz"},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			got := LongestPaliSubstr(test.s)
			want := test.want
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
