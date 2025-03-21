package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func NewBaseObject(img *ebiten.Image, w, h float64) (*BaseObject, error) {
	return &BaseObject{
		image:   img,
		xpos:    0,
		ypos:    0,
		width:   w,
		height:  h,
		visible: true,
	}, nil
}

type BaseObject struct {
	image   *ebiten.Image
	xpos    float64
	ypos    float64
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
	op.GeoM.Translate(o.xpos, o.ypos)
	screen.DrawImage(o.image, op)
}

func (o *BaseObject) SetPosition(x, y float64) {
	o.xpos = x
	o.ypos = y
}

func (o *BaseObject) CenterPosition() (float64, float64) {
	return o.xpos + o.width/2, o.ypos + o.height/2
}

func (o *BaseObject) Distance(p *Player) float64 {
	x1, y1 := p.CenterPosition()
	x2, y2 := o.CenterPosition()
	return Distance(x1, y1, x2, y2)
}

func (o *BaseObject) Hide() {
	o.visible = false
}

func (o *BaseObject) IsVisible() bool {
	return o.visible
}
