package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	tileSet      *ebiten.Image
	turn         int
	heroTurn     bool
	hero         Entity
	entitiesList []Entity
	levelMap     []Cell
	levelWidth   int
	levelHeight  int
	viewWidth    int
	viewHeight   int
}

func NewGame() *Game {
	img, _, err := ebitenutil.NewImageFromFile("./assets/df16x16.png")

	if err != nil {
		log.Fatal("Game - Error when opening file: ", err)
	}

	drunkBot := NewEntity(20, 20, int('b'), DrunkBot)

	width := 1280 / 16
	height := 720 / 16
	//level := make([]Cell, 0)
	//for i := 0; i < width*height; i++ {
	//	level = append(level, *NewCell(i%width, i/width, Wall))
	//}

	level := randomMap(width, height)

	list := make([]Entity, 0)
	list = append(list, *drunkBot)
	return &Game{
		tileSet:      img,
		heroTurn:     true,
		hero:         *NewEntity(0, 0, 64, Hero),
		entitiesList: list,
		turn:         0,
		viewWidth:    1280,
		viewHeight:   720,
		levelMap:     level,
		levelWidth:   1280 / 16,
		levelHeight:  720 / 16,
	}
}

func (g *Game) Update() error {
	// Hero turn
	if g.heroTurn {
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			g.hero.dirX = -1
			g.heroTurn = false

		} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			g.hero.dirX = 1
			g.heroTurn = false

		} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			g.hero.dirY = -1
			g.heroTurn = false

		} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			g.hero.dirY = 1
			g.heroTurn = false

		}
		if !checkDestination(g.hero.x+g.hero.dirX, g.hero.y+g.hero.dirY, &g.levelMap, g.levelWidth) {
			g.hero.dirX = 0
			g.hero.dirY = 0
		}

		g.hero.UpdatePosition()

	} else {
		g.turn += 1
		// Entities turn

		g.heroTurn = true

	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Turn: %d - TPS: %0.2f", g.turn, ebiten.CurrentTPS()))
	for i := 0; i < 3600; i++ {
		op := &ebiten.DrawImageOptions{}
		opBg := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((i%80)*16), float64((i/80)*16)+16)
		op.ColorM.Scale(rand.Float64(), 0, 1.0, 1.0)

		opBg.GeoM.Translate(float64((i%80)*16), float64((i/80)*16)+16)
		opBg.ColorM.Scale(0.0, rand.Float64(), rand.Float64(), 1.0)
		sx := (g.levelMap[i].glyph % 16) * 16
		sy := (g.levelMap[i].glyph / 16) * 16
		//screen.DrawImage(g.tileSet.SubImage(image.Rect(11*16, 13*16, 12*16, 14*16)).(*ebiten.Image), opBg)
		screen.DrawImage(g.tileSet.SubImage(image.Rect(sx, sy, sx+16, sy+16)).(*ebiten.Image), op)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.hero.x*16), float64(g.hero.y*16+16))
	sx := (g.hero.glyph % 16) * 16
	sy := (g.hero.glyph / 16) * 16
	screen.DrawImage(g.tileSet.SubImage(image.Rect(sx, sy, sx+16, sy+16)).(*ebiten.Image), op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.viewWidth, g.viewHeight
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Grogue")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
