package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func NewBaseObject(img *ebiten.Image, w, h float64) (*BaseObject, error) {
	return &BaseObject{
		image: img,
		pos: Point{
			X: 0,
			Y: 0,
		},
		width:   w,
		height:  h,
		visible: true,
	}, nil
}

type BaseObject struct {
	image   *ebiten.Image
	pos     Point
	width   float64
	height  float64
	visible bool
}

func (o *BaseObject) Draw(screen *ebiten.Image) {
	if !o.visible {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(o.pos.X, o.pos.Y)
	screen.DrawImage(o.image, op)
}

func (o *BaseObject) SetPosition(x, y float64) {
	o.pos.X = x
	o.pos.Y = y
}

func (o *BaseObject) CenterPosition() Point {
	return Point{
		X: o.pos.X + o.width/2,
		Y: o.pos.Y + o.height/2,
	}
}

func (o *BaseObject) Distance(p *Player) float64 {
	start := p.CenterPosition()
	end := o.CenterPosition()
	return Distance(start, end)
}

func (o *BaseObject) Hide() {
	o.visible = false
}

func (o *BaseObject) IsVisible() bool {
	return o.visible
}

func (o *BaseObject) Do() {}
