package main

type EntityType int64

const (
	Hero EntityType = iota
	DrunkBot
)

type Entity struct {
	x          int
	y          int
	glyph      int
	entityType EntityType
}

func NewEntity(x, y, glyph int, entityType EntityType) *Entity {

	return &Entity{
		x:          x,
		y:          y,
		glyph:      glyph,
		entityType: entityType,
	}
}
