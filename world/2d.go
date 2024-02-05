package world

import (
	"embed"

	"github.com/808bitt/open-world/entity"
	"github.com/808bitt/open-world/input"
	"github.com/808bitt/open-world/tilemap"
	"github.com/808bitt/open-world/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type World2D struct {
	Width    int
	Height   int
	GridSize int
	TileMap  *tilemap.TileMap
}

func NewWorld2D(width, height, gridSize int, assets *embed.FS) *World2D {
	return &World2D{
		Width:    width / gridSize,
		Height:   height / gridSize,
		GridSize: gridSize,
		TileMap:  tilemap.NewTileMap(width, height, gridSize, assets),
	}
}

func (w *World2D) Draw(screen *ebiten.Image, timeOfDay int) {
	w.TileMap.Draw(screen)
}

func (w *World2D) Update(mouse *input.MouseInput, player *entity.Player) {
	player.Walking = false
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		player.Direction = util.Up.Int()
		if player.Y-player.MoveSpeed >= 0 && w.TileMap.Tiles[player.X/w.GridSize][(player.Y-player.MoveSpeed)/w.GridSize].Walkable {
			player.WalkUp()
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		player.Direction = util.Down.Int()
		if player.Y+player.MoveSpeed < w.Height*w.GridSize && w.TileMap.Tiles[player.X/w.GridSize][(player.Y+player.MoveSpeed)/w.GridSize].Walkable {
			player.WalkDown()
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		player.Direction = util.Left.Int()
		if player.X-player.MoveSpeed >= 0 && w.TileMap.Tiles[(player.X-player.MoveSpeed)/w.GridSize][player.Y/w.GridSize].Walkable {
			player.WalkLeft()
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		player.Direction = util.Right.Int()
		if player.X+player.MoveSpeed < w.Width*w.GridSize && w.TileMap.Tiles[(player.X+player.MoveSpeed)/w.GridSize][player.Y/w.GridSize].Walkable {
			player.WalkRight()
		}
	}
	w.HighlightCursor(mouse, player)

	if player.Walking {
		player.SprIndex++
		if player.SprIndex >= len(player.SprWalkingDown)*10 {
			player.SprIndex = 0
		}
		return
	}
	player.SprIndex = 0

	if mouse.LeftClick {
		pX, pY := w.TileAt(player.X, player.Y)
		mX, mY := w.TileAt(mouse.X, mouse.Y)
		if mX > 0 && mY > 0 && mX < w.Width && mY < w.Height {
			if mX >= pX-2 && mX <= pX+2 && mY >= pY-2 && mY <= pY+2 {
				if mX != pX || mY != pY {
					if w.TileMap.Tiles[mX][mY].Type == tilemap.GrassTile {
						w.TileMap.SetTile(mX, mY, tilemap.NewTile(tilemap.FarmTile, mX, mY, w.TileMap.FarmTile, false))
					} else if w.TileMap.Tiles[mX][mY].Type == tilemap.FarmTile {
						w.TileMap.SetTile(mX, mY, tilemap.NewTile(tilemap.PlantedTile, mX, mY, w.TileMap.PlantedTile, false))
					} else if w.TileMap.Tiles[mX][mY].Type == tilemap.PlantedTile {
						w.TileMap.SetTile(mX, mY, tilemap.NewTile(tilemap.WheatTile, mX, mY, w.TileMap.WheatTile, false))
					}
				}
			}
		}
	}
	if mouse.RightPressed {
		mX, mY := w.TileAt(mouse.X, mouse.Y)
		w.TileMap.SetTile(mX, mY, tilemap.NewTile(tilemap.GrassTile, mX, mY, w.TileMap.GrassTile, true))
	}
}

func (w *World2D) HighlightCursor(mouse *input.MouseInput, player *entity.Player) {
	mX, mY := w.TileAt(mouse.X, mouse.Y)
	pX, pY := w.TileAt(player.X, player.Y)
	for x := 0; x < w.Width; x++ {
		for y := 0; y < w.Height; y++ {
			w.TileMap.Tiles[x][y].Highlight = false
		}
	}
	if mX > 0 && mY > 0 && mX < w.Width && mY < w.Height {
		if mX >= pX-2 && mX <= pX+2 && mY >= pY-2 && mY <= pY+2 {
			w.TileMap.Tiles[mX][mY].Highlight = true
		}
	}
}

func (w *World2D) TileAt(x, y int) (int, int) {
	return x / w.GridSize, y / w.GridSize
}
