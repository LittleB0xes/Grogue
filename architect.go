package main

import "math/rand"

func randomMap(width, height int) []Cell {
	level := make([]Cell, width*height)

	for i := 0; i < width*height; i++ {
		var cellType CellType
		alea := rand.Intn(100)
		if alea < 20 {
			cellType = Wall
		} else {
			cellType = Floor
		}

		x := i % width
		y := i / width

		level[i] = *NewCell(x, y, cellType)
	}

	return level
}
