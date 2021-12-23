package medium

import (
	"reflect"
	"sort"
	"testing"
)

// sortOutput takes the output and sorts each subarray by length of strings
// and then the outer array by len of items
// Example:
// in:	[ ["ab", "a"], ["c"] ]
// out: [ ["c"], ["a", "ab"] ]
func sortOutput(in [][]string) [][]string {
	for _, v := range in {
		// sort subarray by string length asc
		sort.Slice(v, func(i, j int) bool {
			return len(v[i]) < len(v[j])
		})
		// then sort alphabetically
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
	}

	// Finally sort by num items in each subarray
	sort.Slice(in, func(i, j int) bool {
		return len(in[i]) < len(in[j])
	})

	return in
}

func TestGroupAnagrams(t *testing.T) {
	tests := map[string]struct {
		strs []string
		want [][]string
	}{
		"first": {
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			got := sortOutput(GroupAnagrams(test.strs))
			want := test.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
