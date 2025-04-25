package main

type Levels struct {
	list []*Level
	cur  int
}

func NewLevels() *Levels {
	return &Levels{
		list: make([]*Level, 0),
		cur:  0,
	}
}

func (l *Levels) Add(level *Level) {
	l.list = append(l.list, level)
}

func (l *Levels) Current() *Level {
	return l.list[0]
}
