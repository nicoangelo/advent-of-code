package main

type Octopus struct {
	X           int
	Y           int
	EnergyLevel int
	Flashed     bool
}

func (octopus *Octopus) IncrementEnergy() int {
	octopus.EnergyLevel++
	return octopus.EnergyLevel
}

func (octopus *Octopus) FlashIfEnoughEnergy() bool {
	if octopus.EnergyLevel > 9 && !octopus.Flashed {
		octopus.Flashed = true
		return true
	}
	return false
}

func (octopus *Octopus) Reset() {
	if octopus.Flashed {
		octopus.EnergyLevel = 0
		octopus.Flashed = false
	}
}
