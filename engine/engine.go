package engine

import (
	"embed"
	"image/color"
	"log"

	"github.com/808bitt/open-world/entity"
	"github.com/808bitt/open-world/input"
	"github.com/808bitt/open-world/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
		TimeOfDay: 30000,
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
	if e.TimeOfDay > 20000 {
		if e.TimeOfDay <= 30000 {
			shadow := float64(e.TimeOfDay-20000) / 10000 * 150
			flashlight := shadow / 10
			vector.DrawFilledRect(screen, 0, 0, float32(e.World.Width*e.World.GridSize), float32(e.World.Height*e.World.GridSize), color.RGBA{0, 0, 0, uint8(shadow)}, true)
			vector.DrawFilledCircle(screen, float32(e.Player.X), float32(e.Player.Y), 50, color.RGBA{uint8(flashlight), uint8(flashlight), uint8(flashlight), 1}, true)
		} else {
			shadow := float64(40000-e.TimeOfDay) / 10000 * 150
			flashlight := shadow / 10
			vector.DrawFilledRect(screen, 0, 0, float32(e.World.Width*e.World.GridSize), float32(e.World.Height*e.World.GridSize), color.RGBA{0, 0, 0, uint8(shadow)}, true)
			vector.DrawFilledCircle(screen, float32(e.Player.X), float32(e.Player.Y), 50, color.RGBA{uint8(flashlight), uint8(flashlight), uint8(flashlight), 1}, true)
		}
	}
	e.Player.Draw(screen)

	// Draw a simulated light source

	// ebitenutil.DebugPrint(screen, util.Itoa((int(float64(24*e.TimeOfDay/40000))+6)%24)+"h")
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Handle window resize
	return e.World.Width * e.World.GridSize, e.World.Height * e.World.GridSize
}
