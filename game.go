package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

const step = 10

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func NewGopher(left, right, back, front string) (*Gopher, error) {
	leftImage, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("assets/%s", left))
	if err != nil {
		return nil, err
	}

	rightImage, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("assets/%s", right))
	if err != nil {
		return nil, err
	}

	backImage, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("assets/%s", back))
	if err != nil {
		return nil, err
	}

	frontImage, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("assets/%s", front))
	if err != nil {
		return nil, err
	}

	return &Gopher{
		leftImage:  leftImage,
		rightImage: rightImage,
		backImage:  backImage,
		frontImage: frontImage,
		xpos:       0,
		ypos:       0,
		direction:  Right,
	}, nil
}

type Gopher struct {
	leftImage  *ebiten.Image
	rightImage *ebiten.Image
	backImage  *ebiten.Image
	frontImage *ebiten.Image
	xpos       float64
	ypos       float64
	direction  Direction
}

func (g *Gopher) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(g.xpos, g.ypos)
	screen.DrawImage(g.image(), op)
}

func (g *Gopher) MoveLeft() {
	g.direction = Left
	g.xpos -= step
}

func (g *Gopher) MoveRight() {
	g.direction = Right
	g.xpos += step
}

func (g *Gopher) MoveUp() {
	g.direction = Up
	g.ypos -= step
}
func (g *Gopher) MoveDown() {
	g.direction = Down
	g.ypos += step
}

func (g *Gopher) image() *ebiten.Image {
	m := map[Direction]*ebiten.Image{
		Right: g.rightImage,
		Left:  g.leftImage,
		Up:    g.backImage,
		Down:  g.frontImage,
	}
	return m[g.direction]
}

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
