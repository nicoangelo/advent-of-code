package main

type Basin struct {
	Coordinates []*CoordinateValue
}

type BasinList []Basin

func (basin *Basin) TryExpand(system *CoordinateSystem) bool {
	isExpanded := false
	for _, v := range basin.Coordinates {
		neighbors := system.GetNeighborsOf(v)
		for _, nb := range neighbors {
			if nb.Value < 9 && nb.Value > v.Value && basin.AppendCoordinateIfNew(nb) {
				isExpanded = true
			}
		}
	}
	return isExpanded
}

func (basin *Basin) AppendCoordinateIfNew(coordinate *CoordinateValue) bool {
	for _, existing := range basin.Coordinates {
		if existing.X == coordinate.X && existing.Y == coordinate.Y {
			return false
		}
	}
	basin.Coordinates = append(basin.Coordinates, coordinate)
	return true
}
