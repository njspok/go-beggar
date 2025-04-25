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

func (l *Levels) Next() bool {
	if l.cur >= len(l.list)-1 {
		return false
	}

	l.cur++
	return true
}

func (l *Levels) IsFinished() bool {
	for _, lvl := range l.list {
		if !lvl.IsFinish() {
			return false
		}
	}
	return true
}

func (l *Levels) Add(level *Level) {
	l.list = append(l.list, level)
}

func (l *Levels) Current() *Level {
	return l.list[l.cur]
}
