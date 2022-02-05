package easy

// UglyNumber checks if a number's prime factors are limited to 2, 3, and 5
// 1 has no prime factors, therefore all of its prime factors are limited to 2, 3, and 5.
func UglyNumber(n int) bool {
	primes := []int{2, 3, 5}
	for _, p := range primes {
		for n > 0 && n%p == 0 {
			n /= p
		}
	}
	return n == 1
}
