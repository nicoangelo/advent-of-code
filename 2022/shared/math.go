package shared

// Sum adds up all values in the given array and returns the sum
func Sum(values []int) (res int) {
	for _, v := range values {
		res += v
	}
	return res
}
