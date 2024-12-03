package slicemath

// Sum adds up all values in the given array and returns the sum
func Sum(values []int) (res int) {
	for _, v := range values {
		res += v
	}

	return res
}

// Multiply returns the product of all slice values
func Multiply(values []int) (product int) {
	product = 1
	for _, v := range values {
		product *= v
	}
	return product
}
