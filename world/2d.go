package world

import (
	"embed"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type World2D struct {
	Width        int
	Height       int
	GridSize     int
	GrassTile    *ebiten.Image
	WallTop      *ebiten.Image
	WallRight    *ebiten.Image
	WallBot      *ebiten.Image
	WallLeft     *ebiten.Image
	WallTopRight *ebiten.Image
	WallTopLeft  *ebiten.Image
	WallBotRight *ebiten.Image
	WallBotLeft  *ebiten.Image
}

func NewWorld2D(width, height, gridSize int, assets *embed.FS) *World2D {
	return &World2D{
		Width:        width / gridSize,
		Height:       height / gridSize,
		GridSize:     gridSize,
		GrassTile:    LoadImage(assets, "tiles/grass_16.png"),
		WallTop:      LoadImage(assets, "wall/top.png"),
		WallRight:    LoadImage(assets, "wall/right.png"),
		WallBot:      LoadImage(assets, "wall/bottom.png"),
		WallLeft:     LoadImage(assets, "wall/left.png"),
		WallTopRight: LoadImage(assets, "wall/corner_top_right.png"),
		WallTopLeft:  LoadImage(assets, "wall/corner_top_left.png"),
		WallBotRight: LoadImage(assets, "wall/corner_bottom_right.png"),
		WallBotLeft:  LoadImage(assets, "wall/corner_bottom_left.png"),
	}
}

func (w *World2D) Draw(screen *ebiten.Image) {
	// Draw world state
	for x := 0; x < w.Width; x++ {
		for y := 0; y < w.Height; y++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*w.GridSize), float64(y*w.GridSize))
			if x == 1 && y == 1 {
				screen.DrawImage(w.WallTopLeft, op)
			} else if x == w.Width-2 && y == 1 {
				screen.DrawImage(w.WallTopRight, op)
			} else if x == 1 && y == w.Height-2 {
				screen.DrawImage(w.WallBotLeft, op)
			} else if x == w.Width-2 && y == w.Height-2 {
				screen.DrawImage(w.WallBotRight, op)
			} else if x == 1 && y != 0 && y != w.Height-1 {
				screen.DrawImage(w.WallLeft, op)
			} else if x == w.Width-2 && y != 0 && y != w.Height-1 {
				screen.DrawImage(w.WallRight, op)
			} else if y == 1 && x != 0 && x != w.Width-1 {
				screen.DrawImage(w.WallTop, op)
			} else if y == w.Height-2 && x != 0 && x != w.Width-1 {
				screen.DrawImage(w.WallBot, op)
			} else {
				screen.DrawImage(w.GrassTile, op)
			}
		}
	}
}

func (w *World2D) Update() {
	// Update world state
}

func (w *World2D) TileAt(x, y int) (int, int) {
	return x / w.GridSize, y / w.GridSize
}

func LoadImage(assets *embed.FS, path string) *ebiten.Image {
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
