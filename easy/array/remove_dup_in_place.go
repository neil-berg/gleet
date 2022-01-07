package easy

// RemoveDupInPlace removes duplicates in-place and returns the number of
// unique items in the array.
// Eg. [1,1,2,2,3,3] => [1,2,3,_,_,_] k = 3
func RemoveDupInPlace(a []int) int {
	i := 0 // position of last unique item in the final array
	c := 1 // current position in the array

	if len(a) == 0 {
		return 0
	}

	for c < len(a) {
		if a[c] == a[i] {
			// repeat value, advance c until it reaches the next unique item
			for a[c] != a[i]+1 && c < len(a)-1 {
				c++
			}

			if a[c] != a[i] {
				// Swap the next unique element (at c) with pos i and advance i
				a[i+1], a[c] = a[c], a[i+1]
				i++
			}
		} else {
			// unique, advance i
			i++
		}

		c++
	}

	return i + 1
}
