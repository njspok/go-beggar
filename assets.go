package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func NewAssets() *Assets {
	return &Assets{
		images: make(map[string]*ebiten.Image),
	}
}

type Assets struct {
	images map[string]*ebiten.Image
}

func (a *Assets) LoadImage(name string) error {
	img, _, err := ebitenutil.NewImageFromFile(assetFilePath(name))
	if err != nil {
		return err
	}
	a.images[name] = img
	return nil
}

func (a *Assets) LoadImages(names []string) error {
	for _, name := range names {
		if err := a.LoadImage(name); err != nil {
			return err
		}
	}
	return nil
}

func (a *Assets) Image(name string) *ebiten.Image {
	return a.images[name]
}

func assetFilePath(name string) string {
	return fmt.Sprintf("assets/%s", name)
}
