package mutual

// IsEven will check if an integer is even or not by recursively calling IsOdd(n - 1)
func IsEven(n int) bool {
	if n == 0 {
		return true
	}

	// We know if n is even, then n - 1 will be odd
	return IsOdd(n - 1)
}

// IsOdd will check if an integer is odd by checking if the integer is not even
func IsOdd(n int) bool {

	// We know if n is not even, then it will be odd
	return !IsEven(n)
}
