package main

import "github.com/hajimehoshi/ebiten/v2"

func NewRock(image *ebiten.Image, w, h float64) (*Rock, error) {
	obj, err := NewBaseObject(image, w, h)
	if err != nil {
		return nil, err
	}

	return &Rock{
		BaseObject: obj,
	}, nil
}

type Rock struct {
	*BaseObject
}

func (r *Rock) Collision(p *Player) {
	x1, y1 := p.Position()
	w1, h1 := p.Size()

	if IsCollision(x1, y1, w1, h1, r.xpos, r.ypos, r.width, r.height) {
		p.StepBack()
	}
}
