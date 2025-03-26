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
	b.pos.X, b.pos.Y = moveTowards(b.pos, b.target, botStep)

	if b.pos == b.end {
		b.target = b.start
	}

	if b.pos == b.start {
		b.target = b.end
	}
}

func moveTowards(start, end Point, step float64) (float64, float64) {
	dx := end.X - start.X
	dy := end.Y - start.Y

	dist := math.Sqrt(dx*dx + dy*dy)

	if dist <= step {
		return end.X, end.Y
	}

	// вычислаем количество шагов в проекциях
	count := dist / step

	// вычисляем шаги в каждом координате
	stepX := dx / count
	stepY := dy / count

	// Вычисляем новые координаты
	newX := start.X + stepX
	newY := start.Y + stepY

	return newX, newY
}
