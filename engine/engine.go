package engine

import "github.com/hajimehoshi/ebiten/v2"

type Engine struct {
	Width  int
	Height int
}

func NewEngine(width, height int) *Engine {
	return &Engine{width, height}
}

func (e *Engine) Update() error {
	// Update game state
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	// Draw game state
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Handle window resize
	return e.Width, e.Height
}
