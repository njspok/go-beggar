package main

import (
	"bytes"
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/colornames"
	"os"
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
	W    float64
	H    float64
}

type Config struct {
	Width   float64
	Height  float64
	Player  PlayerConfig
	Objects []ObjectConfig
}

type GameStatus int

const (
	GameRunning GameStatus = iota
	GameOver
	GameWin
)

func NewGame(config Config) (*Game, error) {
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowSize(int(config.Width), int(config.Height))

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
			obj, err = NewFood(pos.Type+".png", pos.W, pos.H)
		case "bomb":
			obj, err = NewBomb(pos.Type+".png", pos.W, pos.H)
		case "rock":
			obj, err = NewRock(pos.Type+".png", pos.W, pos.H)
		default:
			return nil, errors.New("invalid type object")
		}

		if err != nil {
			return nil, err
		}

		obj.SetPosition(pos.X, pos.Y)
		objs = append(objs, obj)
	}

	f, err := os.ReadFile("assets/mplus-1p-regular.ttf")
	if err != nil {
		return nil, err
	}

	font, err := text.NewGoTextFaceSource(bytes.NewReader(f))

	g := &Game{
		height: config.Height,
		width:  config.Width,
		keyMap: make(map[ebiten.Key]func()),
		player: player,
		objs:   objs,
		status: GameRunning,
		font:   font,
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
	font   *text.GoTextFaceSource
	status GameStatus
}

func (g *Game) Update() error {
	g.handleKeys()
	g.checkSceneBorders()
	g.checkCollision()
	g.checkGameOver()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	for _, obj := range g.objs {
		obj.Draw(screen)
	}
	g.player.Draw(screen)

	switch g.status {
	case GameRunning:
	case GameOver:
		g.printMessage(screen, "GAME OVER")
	case GameWin:
		g.printMessage(screen, "GAME WIN")
	}
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

func (g *Game) assignKeys() error {
	g.addKeyAction(ebiten.KeyRight, g.player.MoveRight)
	g.addKeyAction(ebiten.KeyLeft, g.player.MoveLeft)
	g.addKeyAction(ebiten.KeyUp, g.player.MoveUp)
	g.addKeyAction(ebiten.KeyDown, g.player.MoveDown)

	return nil
}

func (g *Game) checkGameOver() {
	switch {
	case g.player.IsDied():
		g.status = GameOver
	case g.player.IsSleep():
		g.status = GameWin
	}
}

func (g *Game) printMessage(screen *ebiten.Image, str string) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(g.width/2, g.height/2)
	text.Draw(screen, str, &text.GoTextFace{
		Source: g.font,
		Size:   24,
	}, op)
}
