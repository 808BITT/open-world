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
	IgMap    *util.GridMap
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
		IgMap:    nil,
	}
}

func (e *Engine) Run() {
	if err := ebiten.RunGame(e); err != nil {
		log.Fatal(err)
	}
}

func (e *Engine) Update() error {
	if e.IgMap == nil {
		tileImg := util.LoadImage(e.Assets, "test/tile_top.png")
		tileShape := make([]util.Point, 0)
		for x := 0; x < tileImg.Bounds().Dx(); x++ {
			for y := 0; y < tileImg.Bounds().Dy(); y++ {
				_, _, _, a := tileImg.At(x, y).RGBA()
				if a == 0 {
					continue
				}
				tileShape = append(tileShape, util.Point{X: x, Y: y})
			}
		}

		gridMap := util.NewGridMap(tileShape, e.Settings.Grid.Width, e.Settings.Grid.Height, 64, 64, e.Settings.Screen.Width, e.Assets)
		e.IgMap = gridMap
	}
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	isoGrass := util.LoadImage(e.Assets, "test/iso_grass.png")
	width, height := 64, 64

	for x := 0; x < e.Settings.Grid.Width; x++ {
		for y := 0; y < e.Settings.Grid.Height; y++ {
			opts := &ebiten.DrawImageOptions{}
			xOffset, yOffset := util.GridToIso(x, y, width, height, e.Settings.Screen.Width)
			opts.GeoM.Translate(float64(xOffset), float64(yOffset))
			screen.DrawImage(isoGrass, opts)
		}
	}

	mX, mY := ebiten.CursorPosition()
	lookup := util.IsoPoint{P: util.Point{X: mX, Y: mY}}
	gridPoint, ok := (*e.IgMap.Lookup)[lookup]
	if ok {
		highlight := util.LoadImage(e.Assets, "test/tile_top.png")
		opts := &ebiten.DrawImageOptions{}
		xOffset, yOffset := util.GridToIso(gridPoint.P.X, gridPoint.P.Y, width, height, e.Settings.Screen.Width)
		opts.GeoM.Translate(float64(xOffset), float64(yOffset))
		screen.DrawImage(highlight, opts)
	}
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return e.Settings.Screen.Width, e.Settings.Screen.Height
}
