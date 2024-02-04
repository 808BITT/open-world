package main

import (
	"log"

	"github.com/808bitt/open-world/assets"
	"github.com/808bitt/open-world/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	a := assets.EmbedAssets()
	e := engine.NewEngine(&a)
	if err := ebiten.RunGame(e); err != nil {
		log.Fatal(err)
	}
}
