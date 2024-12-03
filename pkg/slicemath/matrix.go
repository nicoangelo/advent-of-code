package slicemath

type Matrix2D[T comparable] struct {
	values map[Coord2D]T
	size   Coord2D
}

func (m *Matrix2D[T]) Init(size Coord2D) {
	m.values = make(map[Coord2D]T)
	m.size = size
}

func (m *Matrix2D[T]) At(c Coord2D) T {
	v, ok := m.values[c]
	if !ok {
		var empty T
		return empty
	}
	return v
}

func (m *Matrix2D[T]) Set(c Coord2D, v T) {
	if c.X < 0 || c.Y < 0 || c.X > m.size.X-1 || c.Y > m.size.Y-1 {
		return
	}
	m.values[c] = v
}

// SetAndExpand sets the value v at the coordinate c and subsequently
// sets that value to all adjacent coordinates within expandBy length.
// This ignores and overwrites any existing values.
func (m *Matrix2D[T]) SetAndExpand(c Coord2D, v T, expandBy int) {
	if expandBy < 0 {
		panic("Cannot have negative expandBy argument")
	}
	m.values[c] = v
	if expandBy == 0 {
		return
	}
	for i := 1; i <= expandBy; i++ {
		m.Set(Coord2D{X: c.X - i, Y: c.Y}, v)
		m.Set(Coord2D{X: c.X + i, Y: c.Y}, v)
		m.Set(Coord2D{X: c.X, Y: c.Y - i}, v)
		m.Set(Coord2D{X: c.X, Y: c.Y + i}, v)
		m.Set(Coord2D{X: c.X - i, Y: c.Y - i}, v)
		m.Set(Coord2D{X: c.X + i, Y: c.Y + i}, v)
		m.Set(Coord2D{X: c.X + i, Y: c.Y - i}, v)
		m.Set(Coord2D{X: c.X - i, Y: c.Y + i}, v)
	}
}

type Coord2D struct {
	X int
	Y int
}
