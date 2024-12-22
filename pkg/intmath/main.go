package intmath

import "math"

// Abs returns the absolute value of the given integer x
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

// Min returns the smaller of the two integers x and y
func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

// Pow raises x to the power of y
func Pow(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
