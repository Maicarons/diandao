package main

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1024, 576)
	ebiten.SetWindowTitle("♪Music Magnet♪")
	ebiten.SetMaxTPS(30)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	for _, arg := range os.Args {
		switch arg {
		case "-f":
			ebiten.SetFullscreen(true)
		}
	}
	//fmt.Printf("%v,%v,%v", os.Args, Langs, Tr("code"))
	g := &Game{
		scene: &SplashScene{},
	}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
