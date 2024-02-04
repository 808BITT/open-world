package main

import (
	"log"

	"github.com/808bitt/open-world/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	e := engine.NewEngine(800, 600)
	ebiten.SetWindowSize(e.Width, e.Height)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(e); err != nil {
		log.Fatal(err)
	}
}
