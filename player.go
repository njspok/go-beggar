package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	step        = 5
	sleepPoints = 4
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type PlayerStatus int

const (
	Sleeping PlayerStatus = iota
	Awake
	Died
)

func NewPlayer(left, right, back, front, sleep, die *ebiten.Image, w, h float64) (*Player, error) {
	return &Player{
		leftImage:  left,
		rightImage: right,
		backImage:  back,
		frontImage: front,
		sleepImage: sleep,
		dieImage:   die,
		direction:  Right,
		pos: Point{
			X: 0,
			Y: 0,
		},
		prevX:  0,
		prevY:  0,
		width:  w,
		height: h,
		status: Awake,
		points: 0,
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

	pos Point

	prevX float64
	prevY float64

	width  float64
	height float64

	status PlayerStatus

	points int
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(p.pos.X, p.pos.Y)
	screen.DrawImage(p.image(), op)
}

func (p *Player) MoveLeft() {
	if p.isCantMove() {
		return
	}

	p.direction = Left
	p.prevX = p.pos.X
	p.pos.X -= step
}

func (p *Player) MoveRight() {
	if p.isCantMove() {
		return
	}

	p.direction = Right
	p.prevX = p.pos.X
	p.pos.X += step
}

func (p *Player) MoveUp() {
	if p.isCantMove() {
		return
	}

	p.direction = Up
	p.prevY = p.pos.Y
	p.pos.Y -= step
}

func (p *Player) MoveDown() {
	if p.isCantMove() {
		return
	}

	p.direction = Down
	p.prevY = p.pos.Y
	p.pos.Y += step
}

func (p *Player) StepBack() {
	if p.isCantMove() {
		return
	}

	p.pos.X = p.prevX
	p.pos.Y = p.prevY
}

func (p *Player) EndPosition() (x float64, y float64) {
	x = p.pos.X + p.width
	y = p.pos.Y + p.height
	return
}

func (p *Player) Position() (x float64, y float64) {
	x = p.pos.X
	y = p.pos.Y
	return
}

func (p *Player) Size() (w, h float64) {
	return p.width, p.height
}

func (p *Player) CenterPosition() Point {
	return Point{
		X: p.pos.X + p.width/2,
		Y: p.pos.Y + p.height/2,
	}
}

func (p *Player) Width() float64 {
	return p.width
}

func (p *Player) Height() float64 {
	return p.height
}

func (p *Player) Sleep() {
	p.status = Sleeping
}

func (p *Player) Die() {
	p.status = Died
}

func (p *Player) IsSleep() bool {
	return p.status == Sleeping
}

func (p *Player) IsDied() bool {
	return p.status == Died
}

func (p *Player) AddPoint() {
	p.points++
	if p.points == sleepPoints {
		p.Sleep()
	}
}

func (p *Player) image() *ebiten.Image {
	statusImages := map[PlayerStatus]*ebiten.Image{
		Sleeping: p.sleepImage,
		Died:     p.dieImage,
	}

	if i, ok := statusImages[p.status]; ok {
		return i
	}

	directionImages := map[Direction]*ebiten.Image{
		Right: p.rightImage,
		Left:  p.leftImage,
		Up:    p.backImage,
		Down:  p.frontImage,
	}
	return directionImages[p.direction]
}

func (p *Player) isCantMove() bool {
	return p.status == Sleeping || p.status == Died
}
