package main

func NewEat(image string, w, h float64) (*Eat, error) {
	obj, err := NewBaseObject(image, w, h)
	if err != nil {
		return nil, err
	}

	return &Eat{
		BaseObject: obj,
	}, nil
}

type Eat struct {
	*BaseObject
}

func (e *Eat) Collision(p *Player) {
	if e.Distance(p) < 50 {
		e.Hide()
	}
}
