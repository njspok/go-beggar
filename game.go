package main

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

type Object interface {
	Draw(screen *ebiten.Image)
	Collision(player *Player)
	SetPosition(x, y float64)
}

func NewGame(w, h float64) *Game {
	return &Game{
		height: h,
		width:  w,
	}
}

type Game struct {
	player *Player
	objs   []Object
	keyMap map[ebiten.Key]func()
	width  float64
	height float64
}

func (g *Game) Init() error {
	g.keyMap = make(map[ebiten.Key]func())

	player, err := NewPlayer(
		"gopher-left.png",
		"gopher-right.png",
		"gopher-back.png",
		"gopher-front.png",
		"gopher-sleep.png",
		"gopher-die.png",
		playerWidth, playerHeight,
	)
	if err != nil {
		return err
	}
	g.player = player

	// carrots
	objPos := []struct {
		t string
		x float64
		y float64
	}{
		{"carrot", 150, 150},
		{"carrot", 300, 150},
		{"carrot", 0, 300},
		{"bomb", 300, 300},
		{"rock", 300, 50},
	}
	for _, pos := range objPos {
		var obj Object
		var err error

		switch pos.t {
		case "carrot":
			obj, err = NewFood(pos.t+".png", objectWidth, objectHeight)
		case "bomb":
			obj, err = NewBomb(pos.t+".png", objectWidth, objectHeight)
		case "rock":
			obj, err = NewRock(pos.t+".png", objectWidth, objectHeight)
		default:
			return errors.New("invalid type object")
		}

		if err != nil {
			return err
		}

		obj.SetPosition(pos.x, pos.y)
		g.objs = append(g.objs, obj)
	}

	g.addKeyAction(ebiten.KeyRight, g.player.MoveRight)
	g.addKeyAction(ebiten.KeyLeft, g.player.MoveLeft)
	g.addKeyAction(ebiten.KeyUp, g.player.MoveUp)
	g.addKeyAction(ebiten.KeyDown, g.player.MoveDown)

	return nil
}

func (g *Game) Update() error {
	g.handleKeys()
	g.checkSceneBorders()
	g.checkCollision()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	for _, obj := range g.objs {
		obj.Draw(screen)
	}
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.width), int(g.height)
}

func (g *Game) addKeyAction(key ebiten.Key, action func()) {
	g.keyMap[key] = action
}

func (g *Game) handleKeys() {
	for key, action := range g.keyMap {
		if inpututil.KeyPressDuration(key) > 0 {
			action()
		}
	}
}

func (g *Game) checkSceneBorders() {
	x, y := g.player.Position()
	if x < 0 || y < 0 {
		g.player.StepBack()
	}

	ex, ey := g.player.EndPosition()
	if ex > g.width || ey > g.height {
		g.player.StepBack()
	}
}

func (g *Game) checkCollision() {
	for _, obj := range g.objs {
		obj.Collision(g.player)
	}
}
