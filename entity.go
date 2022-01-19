package main

type EntityType int64

const (
	Hero EntityType = iota
	DrunkBot
)

type Entity struct {
	x          int
	y          int
	dirX       int
	dirY       int
	glyph      int
	entityType EntityType
}

func NewEntity(x, y, glyph int, entityType EntityType) *Entity {

	return &Entity{
		x:          x,
		y:          y,
		dirX:       0,
		dirY:       0,
		glyph:      glyph,
		entityType: entityType,
	}
}

func (e *Entity) UpdatePosition() {
	e.x += e.dirX
	e.y += e.dirY

	e.dirX = 0
	e.dirY = 0
}
