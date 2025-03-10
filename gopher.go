package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const step = 10

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func NewGopher(left, right, back, front string, w, h float64) (*Gopher, error) {
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
		direction:  Right,
		xpos:       0,
		ypos:       0,
		width:      w,
		height:     h,
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
	width      float64
	height     float64
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

func (g *Gopher) EndPosition() (x float64, y float64) {
	x = g.xpos + g.width
	y = g.ypos + g.height
	return
}

func (g *Gopher) Position() (x float64, y float64) {
	x = g.xpos
	y = g.ypos
	return
}

func (g *Gopher) SetX(x float64) {
	g.xpos = x
}

func (g *Gopher) SetY(y float64) {
	g.ypos = y
}

func (g *Gopher) Width() float64 {
	return g.width
}

func (g *Gopher) Height() float64 {
	return g.height
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
