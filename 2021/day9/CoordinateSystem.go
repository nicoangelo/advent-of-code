package main

type CoordinateSystem struct {
	Values [][]*CoordinateValue
}

func (system *CoordinateSystem) GetMaxX() int {
	return len(system.Values[0])
}
func (system *CoordinateSystem) GetMaxY() int {
	return len(system.Values)
}

func (system *CoordinateSystem) GetNeighborsOf(coord *CoordinateValue) []*CoordinateValue {
	x := coord.X
	y := coord.Y
	neighbors := make([]*CoordinateValue, 0)
	if x != system.GetMaxX()-1 {
		neighbors = append(neighbors, system.Values[y][x+1])
	}
	if x != 0 {
		neighbors = append(neighbors, system.Values[y][x-1])
	}
	if y != system.GetMaxY()-1 {
		neighbors = append(neighbors, system.Values[y+1][x])
	}
	if y != 0 {
		neighbors = append(neighbors, system.Values[y-1][x])
	}
	return neighbors
}
