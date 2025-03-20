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

type PlayerImagesConfig struct {
	Left  string
	Right string
	Back  string
	Front string
	Sleep string
	Die   string
}

type PlayerConfig struct {
	Width, Height float64
	Images        PlayerImagesConfig
}

type ObjectConfig struct {
	Type string
	X    float64
	Y    float64
}

type Config struct {
	Width   float64
	Height  float64
	Player  PlayerConfig
	Objects []ObjectConfig
}

func NewGame(config Config) (*Game, error) {
	player, err := NewPlayer(
		config.Player.Images.Left,
		config.Player.Images.Right,
		config.Player.Images.Back,
		config.Player.Images.Front,
		config.Player.Images.Sleep,
		config.Player.Images.Die,
		config.Player.Width,
		config.Player.Height,
	)
	if err != nil {
		return nil, err
	}

	var objs []Object
	for _, pos := range config.Objects {
		var obj Object
		var err error

		switch pos.Type {
		case "food":
			obj, err = NewFood(pos.Type+".png", objectWidth, objectHeight)
		case "bomb":
			obj, err = NewBomb(pos.Type+".png", objectWidth, objectHeight)
		case "rock":
			obj, err = NewRock(pos.Type+".png", objectWidth, objectHeight)
		default:
			return nil, errors.New("invalid type object")
		}

		if err != nil {
			return nil, err
		}

		obj.SetPosition(pos.X, pos.Y)
		objs = append(objs, obj)
	}

	g := &Game{
		height: config.Height,
		width:  config.Width,
		keyMap: make(map[ebiten.Key]func()),
		player: player,
		objs:   objs,
	}

	err = g.assignKeys()
	if err != nil {
		return nil, err
	}

	return g, nil
}

type Game struct {
	player *Player
	objs   []Object
	keyMap map[ebiten.Key]func()
	width  float64
	height float64
}

func (g *Game) assignKeys() error {
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
