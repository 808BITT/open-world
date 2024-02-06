package engine

import (
	"log"
	"open-world/assets"
	"open-world/settings"
	"open-world/util"

	"github.com/hajimehoshi/ebiten/v2"
)

type Engine struct {
	Settings *settings.EngineSettings
	Assets   *assets.Assets
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ebiten.SetFullscreen(false)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
}

func NewEngine() *Engine {
	a := assets.EmbedAssets()
	s, e := settings.Load()
	if e != nil {
		log.Fatal(e)
	}

	return &Engine{
		Settings: s,
		Assets:   a,
	}
}

func (e *Engine) Run() {
	if err := ebiten.RunGame(e); err != nil {
		log.Fatal(err)
	}
}

func (e *Engine) Update() error {
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {

	isoGrass := util.LoadImage(e.Assets, "test/iso_grass.png")
	width, height := 64, 64

	for x := 0; x < e.Settings.Screen.Width/width; x++ {
		for y := 0; y < e.Settings.Screen.Width/width; y++ {
			opts := &ebiten.DrawImageOptions{}
			xOffset, yOffset := util.GridToIsometric(x, y, width, height, e.Settings.Screen.Width)
			opts.GeoM.Translate(float64(xOffset), float64(yOffset))
			screen.DrawImage(isoGrass, opts)
		}
	}

	mX, mY := ebiten.CursorPosition()
	x, y := util.IsoToGrid(mX, mY, width, height, e.Settings.Screen.Width)
	log.Println(x, y)

	highlight := util.LoadImage(e.Assets, "test/highlight2.png")
	opts := &ebiten.DrawImageOptions{}
	xOffset, yOffset := util.GridToIsometric(x, y, width, height, e.Settings.Screen.Width)
	opts.GeoM.Translate(float64(xOffset), float64(yOffset))
	screen.DrawImage(highlight, opts)
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return e.Settings.Screen.Width, e.Settings.Screen.Height
}
