package main

type Basin struct {
	Coordinates []*CoordinateValue
}

func (basin *Basin) TryExpand(system *CoordinateSystem) bool {
	return false
}
