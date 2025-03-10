package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

type Game struct {
	gopher *Gopher
	keyMap map[ebiten.Key]func()
}

func (g *Game) Init() error {
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
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	g.gopher.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
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
