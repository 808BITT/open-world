package engine

import (
	"embed"

	"github.com/808bitt/open-world/entity"
	"github.com/808bitt/open-world/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type Engine struct {
	Assets *embed.FS
	World  *world.World2D
	Player *entity.Player
}

func NewEngine(assets *embed.FS) *Engine {
	gridSize := 16
	screenWidth := 1920
	screenHeight := 1080

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Open World")
	ebiten.SetFullscreen(true)
	return &Engine{
		Assets: assets,
		World:  world.NewWorld2D(screenWidth, screenHeight, gridSize, assets),
		Player: entity.NewPlayer(screenWidth/2, screenHeight/2, 1, assets),
	}
}

func (e *Engine) Update() error {
	e.World.Update()         // Update world state
	e.Player.Update(e.World) // Update player state based on world state
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	// Draw game state
	e.World.Draw(screen)
	e.Player.Draw(screen)
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Handle window resize
	return e.World.Width * e.World.GridSize, e.World.Height * e.World.GridSize
}
