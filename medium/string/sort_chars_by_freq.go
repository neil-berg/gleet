package medium

import "sort"

// SortCharsByFreq sorts a string in decreasing order based on the frequency of
// the characters.
// Input: s = "tree"
// Output: "eert"
// Explanation: 'e' appears twice while 'r' and 't' both appear once.
// So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
func SortCharsByFreq(s string) string {
	seen := make(map[string]int)

	for _, r := range s {
		count, exists := seen[string(r)]
		if exists {
			seen[string(r)] = count + 1
		} else {
			seen[string(r)] = 1
		}
	}

	// Sort the map by length
	type Data struct {
		Char  string
		Count int
	}
	var data []Data

	for k, v := range seen {
		data = append(data, Data{Char: k, Count: v})
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Count > data[j].Count
	})

	out := ""
	for _, item := range data {
		charStr := ""
		for i := 0; i < item.Count; i++ {
			charStr += item.Char
		}
		out += charStr
	}
	return out
}
