package main

type CoordinateValue struct {
	X     int
	Y     int
	Value int
}

func (coord *CoordinateValue) IncrementedValueIfLowerThan(others []*CoordinateValue) int {
	for _, v := range others {
		if v.Value <= coord.Value {
			return 0
		}
	}
	return coord.Value + 1
}

func (coord *CoordinateValue) IsLowerThan(others []*CoordinateValue) bool {
	for _, v := range others {
		if v.Value <= coord.Value {
			return false
		}
	}
	return true
}
