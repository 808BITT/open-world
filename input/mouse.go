package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MouseInput struct {
	X, Y                      int
	LeftClick, RightClick     bool
	LeftPressed, RightPressed bool
}

func NewMouseInput() *MouseInput {
	return &MouseInput{}
}

func (m *MouseInput) Update() {
	m.X, m.Y = ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		m.LeftPressed = true
	} else {
		m.LeftPressed = false
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		m.RightPressed = true
	} else {
		m.RightPressed = false
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		m.LeftClick = true
	} else {
		m.LeftClick = false
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		m.RightClick = true
	} else {
		m.RightClick = false
	}
}
