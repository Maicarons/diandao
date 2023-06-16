package main

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type GPlayer struct {
	AudioContext *audio.Context
	bgmPlayer    *audio.Player
	MusicPlay    *audio.Player
}

var Mplay GPlayer

type SceneSwitcher interface {
	ChoScene()
	MainScene()
	PlayScene(n int)
	LoadScene()
	AboutScene()
}

type Scene interface {
	Update(sceneSwitcher SceneSwitcher) error
	Draw(screen *ebiten.Image)
}

type Game struct {
	scene     Scene
	nextScene Scene

	Once sync.Once
}

func (g *Game) Update() error {
	if g.nextScene != nil {
		g.scene = g.nextScene
		g.nextScene = nil
	}
	if err := g.scene.Update(g); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}
