package vectors

import "github.com/nicoangelo/aoc-pkg/intmath"

// Add calculates the sum of vector a and b and returns the result
func Add(a [2]int, b [2]int) [2]int {
	return [2]int{
		a[0] + b[0],
		a[1] + b[1],
	}
}

// Diff calculates the vector difference of a and b
func Diff(a [2]int, b [2]int) [2]int {
	return [2]int{
		a[0] - b[0],
		a[1] - b[1],
	}
}

// MultiplyScalar multiplies every item of the vector
// with the given scalar and returns a new vector
func MultiplyScalar(base [2]int, scalar int) [2]int {
	return [2]int{
		base[0] * scalar,
		base[1] * scalar,
	}
}

// Unity returns the unity vector of the given vector
func Unity(a [2]int) [2]int {
	return [2]int{
		intmath.Sign(a[0]),
		intmath.Sign(a[1]),
	}
}
