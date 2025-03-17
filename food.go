package main

func NewFood(image string, w, h float64) (*Food, error) {
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
	if e.Distance(p) < 50 {
		if e.IsVisible() {
			e.Hide()
			p.AddPoint()
		}
	}
}
