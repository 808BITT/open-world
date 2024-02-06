package util

import (
	"image/png"
	"log"
	"open-world/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadImage(assets *assets.Assets, path string) *ebiten.Image {
	f, err := assets.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Fatalf("Failed to decode image file: %v", err)
	}

	image := ebiten.NewImageFromImage(img)
	return image
}
