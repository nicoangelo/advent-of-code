package day12

import "github.com/nicoangelo/aoc-pkg/slicemath"

type FieldCollection struct {
	fields      map[int]*Field
	curFieldKey int
}

type Field struct {
	tiles         map[slicemath.Coord2D]bool
	area          int
	circumference int
}

func (fc *FieldCollection) AddTile(matrix *slicemath.Matrix2D[rune], curTile slicemath.Coord2D) {

	addField := true
	neighbours := 0

	neighbourkeys := []int{}
	fieldKey := 0

	for _, dir := range directions {

		neighbour := curTile.Add(dir)

		if !matrix.IsOutOfBounds(neighbour) && (matrix.At(neighbour) == matrix.At(curTile)) {
			neighbours++
			for key, field := range fc.fields {
				if field.tiles[neighbour] {
					addField = false
					neighbourkeys = append(neighbourkeys, key)
				}
			}
		}
	}

	if addField {
		newField := &Field{tiles: map[slicemath.Coord2D]bool{}, area: 0, circumference: 0}
		fc.curFieldKey++
		fc.fields[fc.curFieldKey] = newField
		fieldKey = fc.curFieldKey
	} else {
		fieldKey = fc.MergeFields(neighbourkeys)
	}

	curField := fc.fields[fieldKey]

	curField.tiles[curTile] = true
	curField.area += 1
	curField.circumference += (4 - neighbours)

}

func (fc *FieldCollection) MergeFields(fieldKeys []int) int {

	first := fieldKeys[0]

	for _, v := range fieldKeys[1:] {
		if v != first {
			if _, ok := fc.fields[v]; ok {
				for i, w := range fc.fields[v].tiles {
					fc.fields[first].tiles[i] = w
				}
				fc.fields[first].area += fc.fields[v].area
				fc.fields[first].circumference += fc.fields[v].circumference
				delete(fc.fields, v)
			}
		}
	}
	return first
}

var directions []slicemath.Coord2D = []slicemath.Coord2D{
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
}
