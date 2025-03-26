package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

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
		target:     end,
	}

	b.SetPosition(start.X, start.Y)

	return b, nil
}

type Bot struct {
	*BaseObject
	start  Point
	end    Point
	target Point
}

func (b *Bot) Collision(p *Player) {
	if b.Distance(p) < collisionDistance {
		p.Die()
	}
}

func (b *Bot) Do() {
	b.xpos, b.ypos = moveTowards(b.xpos, b.ypos, b.target.X, b.target.Y, botStep)

	if b.xpos == b.end.X && b.ypos == b.end.Y {
		b.target = b.start
	}

	if b.xpos == b.start.X && b.ypos == b.start.Y {
		b.target = b.end
	}
}

func moveTowards(x1, y1, x2, y2, d float64) (float64, float64) {
	dx := x2 - x1
	dy := y2 - y1

	dist := math.Sqrt(dx*dx + dy*dy)

	if dist <= d {
		return x2, y2
	}

	// вычислаем количество шагов в проекциях
	count := dist / d

	// вычисляем шаги в каждом координате
	stepX := dx / count
	stepY := dy / count

	// Вычисляем новые координаты
	newX := x1 + stepX
	newY := y1 + stepY

	return newX, newY
}
