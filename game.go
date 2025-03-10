package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

type Game struct {
	gopher *Gopher
}

func (g *Game) Init() error {
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

	return nil
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.gopher.MoveRight()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.gopher.MoveLeft()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.gopher.MoveUp()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.gopher.MoveDown()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	g.gopher.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
