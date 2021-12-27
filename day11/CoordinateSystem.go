package main

type CoordinateSystem struct {
	Octopuses [][]*Octopus
}

func (system *CoordinateSystem) GetMaxX() int {
	return len(system.Octopuses[0]) - 1
}
func (system *CoordinateSystem) GetMaxY() int {
	return len(system.Octopuses) - 1
}

func (system *CoordinateSystem) GetNeighborsOf(coord *Octopus) []*Octopus {
	x := coord.X
	y := coord.Y
	neighbors := make([]*Octopus, 0)
	if x != system.GetMaxX() {
		neighbors = append(neighbors, system.Octopuses[y][x+1])
		if y != 0 {
			neighbors = append(neighbors, system.Octopuses[y-1][x+1])
		}
		if y != system.GetMaxY() {
			neighbors = append(neighbors, system.Octopuses[y+1][x+1])
		}
	}
	if x != 0 {
		neighbors = append(neighbors, system.Octopuses[y][x-1])
		if y != 0 {
			neighbors = append(neighbors, system.Octopuses[y-1][x-1])
		}
		if y != system.GetMaxY() {
			neighbors = append(neighbors, system.Octopuses[y+1][x-1])
		}
	}
	if y != system.GetMaxY() {
		neighbors = append(neighbors, system.Octopuses[y+1][x])
	}
	if y != 0 {
		neighbors = append(neighbors, system.Octopuses[y-1][x])
	}
	return neighbors
}

func (system *CoordinateSystem) Step() int {
	totalFlashes := 0
	system.increaseEnergyLevels()
	for {
		flashes := system.flashIfEnoughEnergy()
		totalFlashes += flashes
		if flashes == 0 {
			break
		}
	}
	system.resetEnergyLevels()
	return totalFlashes
}

func (system *CoordinateSystem) increaseEnergyLevels() {
	for _, row := range system.Octopuses {
		for _, col := range row {
			col.IncrementEnergy()
		}
	}
}

func (system *CoordinateSystem) flashIfEnoughEnergy() int {
	flashes := 0
	for _, row := range system.Octopuses {
		for _, col := range row {
			if col.FlashIfEnoughEnergy() {
				flashes++
				neighbors := system.GetNeighborsOf(col)
				for _, nb := range neighbors {
					nb.IncrementEnergy()
				}
			}
		}
	}
	return flashes
}

func (system *CoordinateSystem) resetEnergyLevels() {
	for _, row := range system.Octopuses {
		for _, col := range row {
			col.Reset()
		}
	}
}
