package main

type Entity struct {
	x     int
	y     int
	glyph int
}

func NewEntity(x, y, glyph int) *Entity {

	return &Entity{
		x:     x,
		y:     y,
		glyph: glyph,
	}
}
