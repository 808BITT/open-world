package engine

import (
	"embed"
	"log"

	"github.com/808bitt/open-world/entity"
	"github.com/808bitt/open-world/input"
	"github.com/808bitt/open-world/util"
	"github.com/808bitt/open-world/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Engine struct {
	Assets    *embed.FS
	World     *world.World2D
	Player    *entity.Player
	Mouse     *input.MouseInput
	TimeOfDay int
}

func NewEngine(assets *embed.FS) *Engine {
	gridSize := 16
	screenWidth := 1920 / 2
	screenHeight := 1080 / 2

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Open World")
	ebiten.SetFullscreen(true)
	return &Engine{
		Assets:    assets,
		World:     world.NewWorld2D(screenWidth, screenHeight, gridSize, assets),
		Player:    entity.NewPlayer(screenWidth/2, screenHeight/2, 1, assets),
		Mouse:     input.NewMouseInput(),
		TimeOfDay: 0,
	}
}

func (e *Engine) Update() error {
	e.Mouse.Update()                  // Update mouse state
	e.Player.Update()                 // Update player state based on world state
	e.World.Update(e.Mouse, e.Player) // Update world state based on mouse state
	if e.TimeOfDay == 0 {
		log.Println("Morning")
	}
	if e.TimeOfDay == 10000 {
		log.Println("Noon")
	}
	if e.TimeOfDay == 20000 {
		log.Println("Night")
	}
	if e.TimeOfDay == 30000 {
		log.Println("Midnight")
	}

	e.TimeOfDay++        // Update day timer
	e.TimeOfDay %= 40000 // Reset day timer
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	// Draw game state
	e.World.Draw(screen, e.TimeOfDay)
	e.Player.Draw(screen)
	ebitenutil.DebugPrint(screen, util.Itoa(int(float64(24*e.TimeOfDay/40000)))+"h")
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Handle window resize
	return e.World.Width * e.World.GridSize, e.World.Height * e.World.GridSize
}
