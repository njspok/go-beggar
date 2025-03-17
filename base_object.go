package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math"
)

func NewBaseObject(image string, w, h float64) (*BaseObject, error) {
	img, _, err := ebitenutil.NewImageFromFile(imagePath(image))
	if err != nil {
		return nil, err
	}

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

func Distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

func IsCollision(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
	if x1+w1 < x2 || x2+w2 < x1 {
		return false
	}
	if y1+h1 < y2 || y2+h2 < y1 {
		return false
	}
	return true
}
