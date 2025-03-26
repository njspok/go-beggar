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
	ppos := p.Position()
	w1, h1 := p.Size()

	if IsCollision(ppos.X, ppos.Y, w1, h1, r.pos.X, r.pos.Y, r.width, r.height) {
		p.StepBack()
	}
}
