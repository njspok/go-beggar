package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

type Object interface {
	Draw(screen *ebiten.Image)
	Collision(player *Player)
}

type Game struct {
	player *Player
	objs   []Object
	keyMap map[ebiten.Key]func()
	width  float64
	height float64
}

func (g *Game) Init() error {
	g.width = 640
	g.height = 480

	g.keyMap = make(map[ebiten.Key]func())

	player, err := NewPlayer(
		"gopher-left.png",
		"gopher-right.png",
		"gopher-back.png",
		"gopher-front.png",
		128, 128,
	)
	if err != nil {
		return err
	}
	g.player = player

	objPos := []struct {
		x float64
		y float64
	}{
		{150, 150},
		{300, 150},
		{0, 300},
	}
	for _, pos := range objPos {
		obj, err := NewFood("carrot.png", 128, 128)
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

func (g *Game) checkCollision() {
	for _, obj := range g.objs {
		obj.Collision(g.player)
	}
}
