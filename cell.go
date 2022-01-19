package main

type CellType int64

const (
	Floor CellType = iota
	Wall
)

type Cell struct {
	x         int
	y         int
	glyph     int
	cellType  CellType
	crossable bool
}

func NewCell(x, y int, cellType CellType) *Cell {

	var g int
	var c bool

	switch cellType {
	case Floor:
		g = int('.')
		c = true
	case Wall:
		g = int('#')
		c = false

	}
	return &Cell{
		x:         x,
		y:         y,
		glyph:     g,
		crossable: c,
	}
}
