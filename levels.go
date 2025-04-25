package main

type Levels []*Level

func NewLevels() Levels {
	return Levels{}
}

func (l *Levels) Add(level *Level) {
	*l = append(*l, level)
}

func (l *Levels) Current() *Level {
	return (*l)[0]
}
