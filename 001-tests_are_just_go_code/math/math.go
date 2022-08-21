package math

// Sum returns the sum of all the numbers in a slice.
func Sum(numbers []int) int {
	sum := 0

	// Intentional bug.
	for _, n := range numbers {
		sum += n
	}

	return sum
}

// Add returns the sum of intergers a and b.
func Add(a, b int) int {
	return a + b
}
