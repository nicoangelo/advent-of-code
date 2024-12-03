package intmath

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
