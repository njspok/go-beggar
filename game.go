package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

const step = 10

func NewGopher(left, right string) (*Gopher, error) {
	leftImage, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("assets/%s", left))
	if err != nil {
		return nil, err
	}

	rightImage, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("assets/%s", right))
	if err != nil {
		return nil, err
	}

	return &Gopher{
		leftImage:  leftImage,
		rightImage: rightImage,
		xpos:       0,
		ypos:       0,
	}, nil
}

type Gopher struct {
	leftImage  *ebiten.Image
	rightImage *ebiten.Image
	xpos       float64
	ypos       float64
}

func (g *Gopher) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(g.xpos, g.ypos)
	screen.DrawImage(g.leftImage, op)
}

func (g *Gopher) MoveLeft() {
	g.xpos -= step
}

func (g *Gopher) MoveRight() {
	g.xpos += step
}

func (g *Gopher) MoveUp() {
	g.ypos -= step
}
func (g *Gopher) MoveDown() {
	g.ypos += step
}

type Game struct {
	gopher *Gopher
}

func (g *Game) Init() error {
	gopher, err := NewGopher("gopher-left.png", "gopher-right.png")
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
