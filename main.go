package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1024
	screenHeight = 768

	playerWidth  = 128
	playerHeight = 128

	objectWidth  = 128
	objectHeight = 128
)

func main() {
	game, err := NewGame(Config{
		Width:  screenWidth,
		Height: screenHeight,
		Player: PlayerConfig{
			Width:  playerWidth,
			Height: playerHeight,
			Images: PlayerImagesConfig{
				Left:  "gopher-left.png",
				Right: "gopher-right.png",
				Back:  "gopher-back.png",
				Front: "gopher-front.png",
				Sleep: "gopher-sleep.png",
				Die:   "gopher-die.png",
			},
		},
		Levels: []LevelConfig{
			{
				Player: PlayerLevelConfig{
					X: 0,
					Y: 0,
				},
				Objects: []any{
					FoodConfig{X: 500, Y: 0, W: objectWidth, H: objectHeight},
				},
			},
			{
				Player: PlayerLevelConfig{
					X: 0,
					Y: 0,
				},
				Objects: []any{
					FoodConfig{X: 500, Y: 0, W: objectWidth, H: objectHeight},
					FoodConfig{X: 0, Y: 300, W: objectWidth, H: objectHeight},
					FoodConfig{X: 900, Y: 300, W: objectWidth, H: objectHeight},
					FoodConfig{X: 300, Y: 600, W: objectWidth, H: objectHeight},

					BombConfig{X: 500, Y: 300, W: objectWidth, H: objectHeight},

					RockConfig{X: 100, Y: 150, W: objectWidth, H: objectHeight},
					RockConfig{X: 800, Y: 150, W: objectWidth, H: objectHeight},
					RockConfig{X: 100, Y: 500, W: objectWidth, H: objectHeight},
					RockConfig{X: 800, Y: 550, W: objectWidth, H: objectHeight},

					BotConfig{
						W:      objectWidth,
						H:      objectHeight,
						StartX: 200,
						StartY: 150,
						EndX:   600,
						EndY:   150,
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
