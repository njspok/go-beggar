package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

type Game struct {
	gopher *Gopher
	keyMap map[ebiten.Key]func()
	width  int
	height int
}

func (g *Game) Init() error {
	g.width = 320
	g.height = 240

	g.keyMap = make(map[ebiten.Key]func())

	gopher, err := NewGopher(
		"gopher-left.png",
		"gopher-right.png",
		"gopher-back.png",
		"gopher-front.png",
	)
	if err != nil {
		return err
	}

	g.gopher = gopher

	g.addKeyAction(ebiten.KeyRight, g.gopher.MoveRight)
	g.addKeyAction(ebiten.KeyLeft, g.gopher.MoveLeft)
	g.addKeyAction(ebiten.KeyUp, g.gopher.MoveUp)
	g.addKeyAction(ebiten.KeyDown, g.gopher.MoveDown)

	return nil
}

func (g *Game) Update() error {
	g.handleKeys()
	g.checkBorders()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	g.gopher.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}

func (g *Game) addKeyAction(key ebiten.Key, action func()) {
	g.keyMap[key] = action
}

func (g *Game) handleKeys() {
	for key, action := range g.keyMap {
		if inpututil.IsKeyJustPressed(key) {
			action()
		}
	}
}

func (g *Game) checkBorders() {
	x, y := g.gopher.Position()
	if x <= 0 {
		g.gopher.SetX(0)
	}

	if y <= 0 {
		g.gopher.SetY(0)
	}
}
