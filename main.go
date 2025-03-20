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

	collisionDistance = 50
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
		Objects: []ObjectConfig{
			{Type: "food", X: 150, Y: 150},
			{Type: "food", X: 300, Y: 150},
			{Type: "food", X: 0, Y: 300},
			{Type: "bomb", X: 300, Y: 300},
			{Type: "rock", X: 300, Y: 50},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
