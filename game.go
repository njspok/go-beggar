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
	Do()
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
	X float64
	Y float64
	W float64
	H float64
}

type FoodConfig ObjectConfig

type RockConfig ObjectConfig

type BombConfig ObjectConfig

type BotConfig struct {
	W      float64
	H      float64
	StartX float64
	StartY float64
	EndX   float64
	EndY   float64
}

type LevelConfig struct {
	Objects []any
}

type Config struct {
	Width  float64
	Height float64
	Player PlayerConfig
	Level  LevelConfig
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

	assets := NewAssets()

	err := assets.LoadImages([]string{
		config.Player.Images.Left,
		config.Player.Images.Right,
		config.Player.Images.Back,
		config.Player.Images.Front,
		config.Player.Images.Sleep,
		config.Player.Images.Die,
	})
	if err != nil {
		return nil, err
	}

	player, err := NewPlayer(
		Point{
			X: 0,
			Y: 0,
		},
		assets.Image(config.Player.Images.Left),
		assets.Image(config.Player.Images.Right),
		assets.Image(config.Player.Images.Back),
		assets.Image(config.Player.Images.Front),
		assets.Image(config.Player.Images.Sleep),
		assets.Image(config.Player.Images.Die),
		config.Player.Width,
		config.Player.Height,
	)
	if err != nil {
		return nil, err
	}

	err = assets.LoadImages([]string{
		"food.png",
		"bomb.png",
		"rock.png",
		"bot.png",
	})
	if err != nil {
		return nil, err
	}

	var objs []Object
	for _, pos := range config.Level.Objects {
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

	font, err := loadFont("mplus-1p-regular.ttf")
	if err != nil {
		return nil, err
	}

	g := &Game{
		assets: assets,
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
	assets *Assets
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
	g.checkGameFinish()
	g.doObjects()

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
	pos := g.player.Position()
	if pos.X < 0 || pos.Y < 0 {
		g.player.StepBack()
	}

	endPos := g.player.EndPosition()
	if endPos.X > g.width || endPos.Y > g.height {
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

func (g *Game) checkGameFinish() {
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

func (g *Game) doObjects() {
	for _, obj := range g.objs {
		obj.Do()
	}
}

func loadFont(file string) (*text.GoTextFaceSource, error) {
	f, err := os.ReadFile(assetFilePath(file))
	if err != nil {
		return nil, err
	}
	return text.NewGoTextFaceSource(bytes.NewReader(f))
}
