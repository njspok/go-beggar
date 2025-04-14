package main

import "errors"

func NewLevel(config LevelConfig, assets *Assets) (*Level, error) {
	var objs []Object
	for _, pos := range config.Objects {
		var obj Object
		var err error

		switch p := pos.(type) {
		case FoodConfig:
			obj, err = NewFood(assets.Image("food.png"), p.W, p.H)
			obj.SetPosition(p.X, p.Y)
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
	}, nil
}

type Level struct {
	Objects []Object
}
