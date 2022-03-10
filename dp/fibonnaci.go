package dp

// Fib returns the nth element in a fibonnaci series
func Fib(n int) int {
	if n <= 2 {
		return 1
	}
	return Fib(n-2) + Fib(n-1)
}
