package entity

import (
	"embed"

	"github.com/808bitt/open-world/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	X, Y             int
	MoveSpeed        int
	Walking          bool
	Direction        int
	SprIndex         int
	SprStandingDown  *ebiten.Image
	SprStandingLeft  *ebiten.Image
	SprStandingRight *ebiten.Image
	SprStandingUp    *ebiten.Image
	SprWalkingDown   []*ebiten.Image
	SprWalkingLeft   []*ebiten.Image
	SprWalkingRight  []*ebiten.Image
	SprWalkingUp     []*ebiten.Image
}

func NewPlayer(x, y, moveSpeed int, assets *embed.FS) *Player {
	return &Player{
		X:                x,
		Y:                y,
		MoveSpeed:        moveSpeed,
		Direction:        util.Down.Int(),
		Walking:          false,
		SprIndex:         0,
		SprStandingDown:  util.LoadImage(assets, "player/standing/down.png"),
		SprStandingLeft:  util.LoadImage(assets, "player/standing/left.png"),
		SprStandingRight: util.LoadImage(assets, "player/standing/right.png"),
		SprStandingUp:    util.LoadImage(assets, "player/standing/up.png"),
		SprWalkingDown: []*ebiten.Image{
			util.LoadImage(assets, "player/walking/down_1.png"),
			util.LoadImage(assets, "player/walking/down_2.png"),
			util.LoadImage(assets, "player/walking/down_3.png"),
			util.LoadImage(assets, "player/walking/down_4.png"),
		},
		SprWalkingLeft: []*ebiten.Image{
			util.LoadImage(assets, "player/walking/left_1.png"),
			util.LoadImage(assets, "player/walking/left_2.png"),
			util.LoadImage(assets, "player/walking/left_3.png"),
			util.LoadImage(assets, "player/walking/left_4.png"),
		},
		SprWalkingRight: []*ebiten.Image{
			util.LoadImage(assets, "player/walking/right_1.png"),
			util.LoadImage(assets, "player/walking/right_2.png"),
			util.LoadImage(assets, "player/walking/right_3.png"),
			util.LoadImage(assets, "player/walking/right_4.png"),
		},
		SprWalkingUp: []*ebiten.Image{
			util.LoadImage(assets, "player/walking/up_1.png"),
			util.LoadImage(assets, "player/walking/up_2.png"),
			util.LoadImage(assets, "player/walking/up_3.png"),
			util.LoadImage(assets, "player/walking/up_4.png"),
		},
	}
}

func (p *Player) Update() {

}

func (p *Player) Draw(screen *ebiten.Image) {
	// Draw player state
	op := &ebiten.DrawImageOptions{}

	if p.Walking {
		i := p.SprIndex / 10
		if p.Direction == util.Up.Int() {
			op.GeoM.Translate(float64(p.X-(p.SprWalkingUp[i].Bounds().Dx()/2)), float64(p.Y-(p.SprWalkingUp[i].Bounds().Dy())))
			screen.DrawImage(p.SprWalkingUp[i], op)
			return
		}
		if p.Direction == util.Down.Int() {
			op.GeoM.Translate(float64(p.X-(p.SprWalkingDown[i].Bounds().Dx()/2)), float64(p.Y-(p.SprWalkingDown[i].Bounds().Dy())))
			screen.DrawImage(p.SprWalkingDown[p.SprIndex/10], op)
			return
		}
		if p.Direction == util.Left.Int() {
			op.GeoM.Translate(float64(p.X-(p.SprWalkingLeft[i].Bounds().Dx()/2)), float64(p.Y-(p.SprWalkingLeft[i].Bounds().Dy())))
			screen.DrawImage(p.SprWalkingLeft[p.SprIndex/10], op)
			return
		}
		if p.Direction == util.Right.Int() {
			op.GeoM.Translate(float64(p.X-(p.SprWalkingRight[i].Bounds().Dx()/2)), float64(p.Y-(p.SprWalkingRight[i].Bounds().Dy())))
			screen.DrawImage(p.SprWalkingRight[p.SprIndex/10], op)
			return
		}
	}
	if p.Direction == util.Up.Int() {
		op.GeoM.Translate(float64(p.X-(p.SprStandingUp.Bounds().Dx()/2)), float64(p.Y-(p.SprStandingUp.Bounds().Dy())))
		screen.DrawImage(p.SprStandingUp, op)
		return
	}
	if p.Direction == util.Down.Int() {
		op.GeoM.Translate(float64(p.X-(p.SprStandingDown.Bounds().Dx()/2)), float64(p.Y-(p.SprStandingDown.Bounds().Dy())))
		screen.DrawImage(p.SprStandingDown, op)
		return
	}
	if p.Direction == util.Left.Int() {
		op.GeoM.Translate(float64(p.X-(p.SprStandingLeft.Bounds().Dx()/2)), float64(p.Y-(p.SprStandingLeft.Bounds().Dy())))
		screen.DrawImage(p.SprStandingLeft, op)
		return
	}
	if p.Direction == util.Right.Int() {
		op.GeoM.Translate(float64(p.X-(p.SprStandingRight.Bounds().Dx()/2)), float64(p.Y-(p.SprStandingRight.Bounds().Dy())))
		screen.DrawImage(p.SprStandingRight, op)
		return
	}
}

func (p *Player) Hitbox() (int, int, int, int) {
	return p.X + 4, p.Y + 26, p.SprStandingDown.Bounds().Dx() - 8, p.SprStandingDown.Bounds().Dy() - 26
}

func (p *Player) Move(dx, dy int) {
	p.X += dx * p.MoveSpeed
	p.Y += dy * p.MoveSpeed
}

func (p *Player) WalkUp() {
	p.Walking = true
	p.Move(0, -1)
}

func (p *Player) WalkDown() {
	p.Walking = true
	p.Move(0, 1)
}

func (p *Player) WalkLeft() {
	p.Walking = true
	p.Move(-1, 0)
}

func (p *Player) WalkRight() {
	p.Walking = true
	p.Move(1, 0)
}
