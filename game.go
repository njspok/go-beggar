package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

type Game struct {
	player *Player
	keyMap map[ebiten.Key]func()
	width  float64
	height float64
}

func (g *Game) Init() error {
	g.width = 320
	g.height = 240

	g.keyMap = make(map[ebiten.Key]func())

	gopher, err := NewPlayer(
		"gopher-left.png",
		"gopher-right.png",
		"gopher-back.png",
		"gopher-front.png",
		128, 128,
	)
	if err != nil {
		return err
	}

	g.player = gopher

	g.addKeyAction(ebiten.KeyRight, g.player.MoveRight)
	g.addKeyAction(ebiten.KeyLeft, g.player.MoveLeft)
	g.addKeyAction(ebiten.KeyUp, g.player.MoveUp)
	g.addKeyAction(ebiten.KeyDown, g.player.MoveDown)

	return nil
}

func (g *Game) Update() error {
	g.handleKeys()
	g.checkSceneBorders()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
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
		if inpututil.IsKeyJustPressed(key) {
			action()
		}
	}
}

func (g *Game) checkSceneBorders() {
	x, y := g.player.Position()
	if x <= 0 {
		g.player.SetX(0)
	}
	if y <= 0 {
		g.player.SetY(0)
	}

	ex, ey := g.player.EndPosition()
	if ex >= g.width {
		g.player.SetX(g.width - g.player.Width())
	}
	if ey >= g.height {
		g.player.SetY(g.height - g.player.Height())
	}
}
