package main

import "github.com/hajimehoshi/ebiten/v2"

const collisionDistance = 100

func NewBomb(image *ebiten.Image, w, h float64) (*Bomb, error) {
	obj, err := NewBaseObject(image, w, h)
	if err != nil {
		return nil, err
	}

	return &Bomb{
		BaseObject: obj,
	}, nil
}

type Bomb struct {
	*BaseObject
}

func (e *Bomb) Collision(p *Player) {
	if e.Distance(p) < collisionDistance {
		e.Hide()
		p.Die()
	}
}
