package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math"
)

func NewObject(image string, w, h float64) (*Object, error) {
	img, _, err := ebitenutil.NewImageFromFile(imagePath(image))
	if err != nil {
		return nil, err
	}

	return &Object{
		image:   img,
		xpos:    0,
		ypos:    0,
		width:   w,
		height:  h,
		visible: true,
	}, nil
}

type Object struct {
	image   *ebiten.Image
	xpos    float64
	ypos    float64
	width   float64
	height  float64
	visible bool
}

func (o *Object) Draw(screen *ebiten.Image) {
	if !o.visible {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(o.xpos, o.ypos)
	screen.DrawImage(o.image, op)
}

func (o *Object) SetPosition(x, y float64) {
	o.xpos = x
	o.ypos = y
}

func (o *Object) Collision(p *Player) {
	x1, y1 := p.Position()
	w1 := p.Width()
	h1 := p.Height()

	x2, y2 := o.xpos, o.ypos
	w2 := o.width
	h2 := o.height

	if Distance(x1, y1, w1, h1, x2, y2, w2, h2) < 50 {
		o.visible = false
	}
}

func Distance(x1, y1, w1, h1, x2, y2, w2, h2 float64) float64 {
	cx1 := x1 + w1
	cy1 := y1 + h1
	cx2 := x2 + w2
	cy2 := y2 + h2

	return math.Sqrt(math.Pow(cx1-cx2, 2) + math.Pow(cy1-cy2, 2))
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
