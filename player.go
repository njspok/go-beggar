package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func imagePath(name string) string {
	return fmt.Sprintf("assets/%s", name)
}

const (
	step        = 5
	sleepPoints = 3
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Satus int

const (
	Sleeping Satus = iota
	Awake
	Died
)

func NewPlayer(left, right, back, front, sleep, die string, w, h float64) (*Player, error) {
	leftImage, _, err := ebitenutil.NewImageFromFile(imagePath(left))
	if err != nil {
		return nil, err
	}

	rightImage, _, err := ebitenutil.NewImageFromFile(imagePath(right))
	if err != nil {
		return nil, err
	}

	backImage, _, err := ebitenutil.NewImageFromFile(imagePath(back))
	if err != nil {
		return nil, err
	}

	frontImage, _, err := ebitenutil.NewImageFromFile(imagePath(front))
	if err != nil {
		return nil, err
	}

	sleepImage, _, err := ebitenutil.NewImageFromFile(imagePath(sleep))
	if err != nil {
		return nil, err
	}

	dieImage, _, err := ebitenutil.NewImageFromFile(imagePath(die))
	if err != nil {
		return nil, err
	}

	return &Player{
		leftImage:  leftImage,
		rightImage: rightImage,
		backImage:  backImage,
		frontImage: frontImage,
		sleepImage: sleepImage,
		dieImage:   dieImage,
		direction:  Right,
		xpos:       0,
		ypos:       0,
		width:      w,
		height:     h,
		status:     Awake,
		points:     0,
	}, nil
}

type Player struct {
	Object

	leftImage  *ebiten.Image
	rightImage *ebiten.Image
	backImage  *ebiten.Image
	frontImage *ebiten.Image
	sleepImage *ebiten.Image
	dieImage   *ebiten.Image

	direction Direction

	xpos   float64
	ypos   float64
	width  float64
	height float64

	status Satus

	points int
}

func (g *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(g.xpos, g.ypos)
	screen.DrawImage(g.image(), op)
}

func (g *Player) MoveLeft() {
	if g.isCantMove() {
		return
	}

	g.direction = Left
	g.xpos -= step
}

func (g *Player) MoveRight() {
	if g.isCantMove() {
		return
	}

	g.direction = Right
	g.xpos += step
}

func (g *Player) MoveUp() {
	if g.isCantMove() {
		return
	}

	g.direction = Up
	g.ypos -= step
}

func (g *Player) MoveDown() {
	if g.isCantMove() {
		return
	}

	g.direction = Down
	g.ypos += step
}

func (g *Player) EndPosition() (x float64, y float64) {
	x = g.xpos + g.width
	y = g.ypos + g.height
	return
}

func (g *Player) Position() (x float64, y float64) {
	x = g.xpos
	y = g.ypos
	return
}

func (p *Player) CenterPosition() (x float64, y float64) {
	return p.xpos + p.width/2, p.ypos + p.height/2
}

func (g *Player) SetX(x float64) {
	g.xpos = x
}

func (g *Player) SetY(y float64) {
	g.ypos = y
}

func (g *Player) Width() float64 {
	return g.width
}

func (g *Player) Height() float64 {
	return g.height
}

func (g *Player) Sleep() {
	g.status = Sleeping
}

func (g *Player) Die() {
	g.status = Died
}

func (g *Player) AddPoint() {
	g.points++
	if g.points == sleepPoints {
		g.Sleep()
	}
}

func (g *Player) image() *ebiten.Image {
	statusImages := map[Satus]*ebiten.Image{
		Sleeping: g.sleepImage,
		Died:     g.dieImage,
	}

	if i, ok := statusImages[g.status]; ok {
		return i
	}

	directionImages := map[Direction]*ebiten.Image{
		Right: g.rightImage,
		Left:  g.leftImage,
		Up:    g.backImage,
		Down:  g.frontImage,
	}
	return directionImages[g.direction]
}

func (g *Player) isCantMove() bool {
	return g.status == Sleeping || g.status == Died
}
