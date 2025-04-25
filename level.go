package main

import (
	"errors"
)

func NewLevel(config LevelConfig, assets *Assets, player *Player) (*Level, error) {
	var objs []Object
	var foods int
	for _, pos := range config.Objects {
		var obj Object
		var err error

		switch p := pos.(type) {
		case FoodConfig:
			obj, err = NewFood(assets.Image("food.png"), p.W, p.H)
			obj.SetPosition(p.X, p.Y)
			foods++
		case BombConfig:
			obj, err = NewBomb(assets.Image("bomb.png"), p.W, p.H)
			obj.SetPosition(p.X, p.Y)
		case RockConfig:
			obj, err = NewRock(assets.Image("rock.png"), p.W, p.H)
			obj.SetPosition(p.X, p.Y)
		case BotConfig:
			obj, err = NewBot(
				assets.Image("bot.png"),
				p.W,
				p.H,
				Point{
					X: p.StartX,
					Y: p.StartY,
				},
				Point{
					X: p.EndX,
					Y: p.EndY,
				},
			)
		default:
			return nil, errors.New("invalid type object")
		}

		if err != nil {
			return nil, err
		}

		objs = append(objs, obj)
	}

	return &Level{
		Objects: objs,
		player:  player,
		foods:   foods,
		startPlayerPos: Point{
			X: config.Player.X,
			Y: config.Player.Y,
		},
	}, nil
}

type Level struct {
	Objects        []Object
	player         *Player
	startPlayerPos Point
	foods          int
	isFinished     bool
}

func (l *Level) IsFinish() bool {
	if l.isFinished {
		return true
	}

	if l.player.Score() == l.foods {
		l.isFinished = true
	}

	return l.isFinished
}

func (l *Level) Init() {
	l.player.SetPosition(l.startPlayerPos)
	l.player.SetScore(0)
}
