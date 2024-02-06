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

type Point struct {
	X, Y int
}

func (e *Engine) Update() error {
	// tileImg := util.LoadImage(assets.EmbedAssets(), "test/iso_grass.png")
	// tileShape := make([]Point, 0)
	// for x := 0; x < 64; x++ {
	// 	for y := 0; y < 64; y++ {
	// 		_, _, _, a := tileImg.At(x, y).RGBA()
	// 		if a == 0 {
	// 			continue
	// 		}
	// 		tileShape = append(tileShape, Point{x, y})
	// 	}
	// }

	// // for y := 0; y < 25; y++ {
	// // 	for x := 0; x < 25; x++ {

	// x, y := 0, 0
	// isoX, isoY := util.GridToIso(x, y, 64, 64, 1920)
	// fmt.Println(x, y, isoX, isoY)
	// for _, p := range tileShape {
	// 	fmt.Println("Checking: ", p.X, p.Y)
	// 	// check isoX+p.X, isoY+p.Y to see if it maps back to x, y
	// 	// if it doesnt, panic

	// 	gridX, gridY := util.IsoToGrid(isoX+p.X, isoY+p.Y, 64, 64, 1920)
	// 	if gridX != x || gridY != y {
	// 		log.Println("Mismatch", x, y, gridX, gridY)
	// 		os.Exit(1)
	// 	}
	// }
	// // 	}
	// // }
	//
	// os.Exit(0)
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {

	isoGrass := util.LoadImage(e.Assets, "test/iso_grass.png")
	width, height := 64, 64

	for x := 0; x < 25; x++ {
		for y := 0; y < 25; y++ {
			opts := &ebiten.DrawImageOptions{}
			xOffset, yOffset := util.GridToIso(x, y, width, height, e.Settings.Screen.Width)
			opts.GeoM.Translate(float64(xOffset), float64(yOffset))
			screen.DrawImage(isoGrass, opts)
		}
	}

	mX, mY := ebiten.CursorPosition()
	x, y := util.IsoToGrid(mX, mY, width, height, e.Settings.Screen.Width)

	highlight := util.LoadImage(e.Assets, "test/highlight2.png")
	opts := &ebiten.DrawImageOptions{}
	xOffset, yOffset := util.GridToIso(x, y, width, height, e.Settings.Screen.Width)
	opts.GeoM.Translate(float64(xOffset), float64(yOffset))
	screen.DrawImage(highlight, opts)
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return e.Settings.Screen.Width, e.Settings.Screen.Height
}
