package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
	"log"
)

type Game struct {
	title  string
	gopher *ebiten.Image
	xpos   float64
	run    bool
}

func (g *Game) Init() {
	gopher, _, err := ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	g.gopher = gopher

	g.xpos = 0
	g.title = "Gopher"
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.run = true
	}

	if g.run {
		g.xpos += 3
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.xpos, 10)
	screen.DrawImage(g.gopher, op)

	ebitenutil.DebugPrint(screen, g.title)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
