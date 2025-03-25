package main

import "github.com/hajimehoshi/ebiten/v2"

const botStep = 5

func NewBot(image *ebiten.Image, w, h float64, start, end Point) (*Bot, error) {
	obj, err := NewBaseObject(image, w, h)
	if err != nil {
		return nil, err
	}

	b := &Bot{
		BaseObject: obj,
		start:      start,
		end:        end,
		step:       botStep,
	}

	b.SetPosition(start.X, start.Y)

	return b, nil
}

type Bot struct {
	*BaseObject
	start Point
	end   Point
	step  int
}

func (b *Bot) Collision(p *Player) {
	// todo process collision
}

func (b *Bot) Do() {
	b.xpos += float64(b.step)
	if b.xpos >= b.end.X || b.xpos <= b.end.Y {
		b.step *= -1
	}
}
