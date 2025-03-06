package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
	"log"
	"math/rand"
)

var list = []string{
	"hello, world!",
	"hi, world",
}

type Game struct {
	title  string
	gopher *ebiten.Image
}

func (g *Game) Init() {
	gopher, _, err := ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	g.gopher = gopher
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		n := rand.Intn(len(list))
		g.title = list[n]
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(10, 10)
	screen.DrawImage(g.gopher, op)

	ebitenutil.DebugPrint(screen, g.title)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
