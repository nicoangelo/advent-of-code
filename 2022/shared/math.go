package shared

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

// VectorAdd calculates the sum of vector a and b and returns the result
func VectorAdd(a [2]int, b [2]int) [2]int {
	return [2]int{
		a[0] + b[0],
		a[1] + b[1],
	}
}

// VectorDiff calculates the vector difference of a and b
func VectorDiff(a [2]int, b [2]int) [2]int {
	return [2]int{
		a[0] - b[0],
		a[1] - b[1],
	}
}

// VectorMultiplyScalar multiplies every item of the vector
// with the given scalar and returns a new vector
func VectorMultiplyScalar(base [2]int, scalar int) [2]int {
	return [2]int{
		base[0] * scalar,
		base[1] * scalar,
	}
}

// AbsInt returns the absolute value of the given integer x
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// VectorUnity returns the unity vector of the given vector
func VectorUnity(a [2]int) [2]int {
	return [2]int{
		Sign(a[0]),
		Sign(a[1]),
	}
}

// Sign returns the sign of the given integer x
func Sign(x int) int {
	if x < 0 {
		return -1
	} else if x == 0 {
		return 0
	}
	return 1
}
