package dp

// FibCache is a map of seen nodes and their corresponding values
type FibCache map[int]int

// Fib returns the nth element in a fibonnaci series
func Fib(n int, cache FibCache) int {
	value, exists := cache[n]
	if exists {
		return value
	}

	if n <= 2 {
		return 1
	}

	cache[n] = Fib(n-2, cache) + Fib(n-1, cache)
	return cache[n]
}

// FibTabulation iteratively returns the nth Fib number
func FibTabulation(n int) int {
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1

	for i := 2; i <= n; i++ {
		arr[i] = arr[i-2] + arr[i-1]
	}
	return arr[n]
}
