package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	step = 5
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

func NewPlayer(
	pos Point,
	left, right, back, front, sleep, die *ebiten.Image,
	w, h float64,
) (*Player, error) {
	return &Player{
		leftImage:  left,
		rightImage: right,
		backImage:  back,
		frontImage: front,
		sleepImage: sleep,
		dieImage:   die,
		direction:  Right,
		pos:        pos,
		prev:       pos,
		width:      w,
		height:     h,
		status:     Awake,
		score:      0,
	}, nil
}

type Player struct {
	leftImage  *ebiten.Image
	rightImage *ebiten.Image
	backImage  *ebiten.Image
	frontImage *ebiten.Image
	sleepImage *ebiten.Image
	dieImage   *ebiten.Image

	direction Direction

	pos  Point
	prev Point

	width  float64
	height float64

	status PlayerStatus

	score int
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
	p.prev.X = p.pos.X
	p.pos.X -= step
}

func (p *Player) MoveRight() {
	if p.isCantMove() {
		return
	}

	p.direction = Right
	p.prev.X = p.pos.X
	p.pos.X += step
}

func (p *Player) MoveUp() {
	if p.isCantMove() {
		return
	}

	p.direction = Up
	p.prev.Y = p.pos.Y
	p.pos.Y -= step
}

func (p *Player) MoveDown() {
	if p.isCantMove() {
		return
	}

	p.direction = Down
	p.prev.Y = p.pos.Y
	p.pos.Y += step
}

func (p *Player) StepBack() {
	if p.isCantMove() {
		return
	}

	p.pos.X = p.prev.X
	p.pos.Y = p.prev.Y
}

func (p *Player) EndPosition() Point {
	return Point{
		X: p.pos.X + p.width,
		Y: p.pos.Y + p.height,
	}
}

func (p *Player) SetPosition(pos Point) {
	p.pos = pos
}

func (p *Player) Position() Point {
	return p.pos
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

func (p *Player) IncScore() {
	p.score++
}

func (p *Player) Score() int {
	return p.score
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

func (p *Player) SetScore(score int) {
	p.score = score
}
