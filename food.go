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

func (f *Food) Collision(p *Player) {
	if f.Distance(p) < collisionDistance {
		if f.IsVisible() {
			f.Hide()
			p.IncScore()
		}
	}
}
