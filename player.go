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
		xpos:       0,
		ypos:       0,
		prevX:      0,
		prevY:      0,
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

	xpos  float64
	ypos  float64
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
	op.GeoM.Translate(p.xpos, p.ypos)
	screen.DrawImage(p.image(), op)
}

func (p *Player) MoveLeft() {
	if p.isCantMove() {
		return
	}

	p.direction = Left
	p.prevX = p.xpos
	p.xpos -= step
}

func (p *Player) MoveRight() {
	if p.isCantMove() {
		return
	}

	p.direction = Right
	p.prevX = p.xpos
	p.xpos += step
}

func (p *Player) MoveUp() {
	if p.isCantMove() {
		return
	}

	p.direction = Up
	p.prevY = p.ypos
	p.ypos -= step
}

func (p *Player) MoveDown() {
	if p.isCantMove() {
		return
	}

	p.direction = Down
	p.prevY = p.ypos
	p.ypos += step
}

func (p *Player) StepBack() {
	if p.isCantMove() {
		return
	}

	p.xpos = p.prevX
	p.ypos = p.prevY
}

func (p *Player) EndPosition() (x float64, y float64) {
	x = p.xpos + p.width
	y = p.ypos + p.height
	return
}

func (p *Player) Position() (x float64, y float64) {
	x = p.xpos
	y = p.ypos
	return
}

func (p *Player) Size() (w, h float64) {
	return p.width, p.height
}

func (p *Player) CenterPosition() (x float64, y float64) {
	return p.xpos + p.width/2, p.ypos + p.height/2
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
