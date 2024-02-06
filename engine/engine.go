package engine

import (
	"log"
	"open-world/assets"
	"open-world/settings"

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
	// draw an isometric grid
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return e.Settings.Screen.Width, e.Settings.Screen.Height
}
