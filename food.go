package main

import "github.com/hajimehoshi/ebiten/v2"

func NewFood(image *ebiten.Image, w, h float64) (*Food, error) {
	obj, err := NewBaseObject(image, w, h)
	if err != nil {
		return nil, err
	}

	return &Food{
		BaseObject: obj,
	}, nil
}

type Food struct {
	*BaseObject
}

func (e *Food) Collision(p *Player) {
	if e.Distance(p) < collisionDistance {
		if e.IsVisible() {
			e.Hide()
			p.AddPoint()
		}
	}
}
