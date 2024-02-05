package tilemap

import (
	"embed"
	"image/color"

	"github.com/808bitt/open-world/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type TileMap struct {
	Width, Height int
	GridSize      int
	Tiles         [][]*Tile
	GrassTile     *ebiten.Image
	FarmTile      *ebiten.Image
	PlantedTile   *ebiten.Image
	WheatTile     *ebiten.Image
	WallTop       *ebiten.Image
	WallRight     *ebiten.Image
	WallBot       *ebiten.Image
	WallLeft      *ebiten.Image
	WallTopRight  *ebiten.Image
	WallTopLeft   *ebiten.Image
	WallBotRight  *ebiten.Image
	WallBotLeft   *ebiten.Image
}

func NewTileMap(width, height, gridSize int, assets *embed.FS) *TileMap {
	tiles := make([][]*Tile, width/gridSize)
	for i := range tiles {
		tiles[i] = make([]*Tile, height/gridSize)
	}

	for x := 0; x < width/gridSize; x++ {
		for y := 0; y < height/gridSize; y++ {
			tiles[x][y] = NewTile(GrassTile, x, y, util.LoadImage(assets, "tile/grass.png"), true)
		}
	}

	return &TileMap{
		Width:        width,
		Height:       height,
		GridSize:     gridSize,
		Tiles:        tiles,
		GrassTile:    util.LoadImage(assets, "tile/grass.png"),
		FarmTile:     util.LoadImage(assets, "tile/farmland.png"),
		PlantedTile:  util.LoadImage(assets, "tile/farmland_planted.png"),
		WheatTile:    util.LoadImage(assets, "tile/farmland_wheat.png"),
		WallTop:      util.LoadImage(assets, "wall/top.png"),
		WallRight:    util.LoadImage(assets, "wall/right.png"),
		WallBot:      util.LoadImage(assets, "wall/bottom.png"),
		WallLeft:     util.LoadImage(assets, "wall/left.png"),
		WallTopRight: util.LoadImage(assets, "wall/corner_top_right.png"),
		WallTopLeft:  util.LoadImage(assets, "wall/corner_top_left.png"),
		WallBotRight: util.LoadImage(assets, "wall/corner_bottom_right.png"),
		WallBotLeft:  util.LoadImage(assets, "wall/corner_bottom_left.png"),
	}
}

func (t *TileMap) SetTile(x, y int, tile *Tile) {
	t.Tiles[x][y] = tile
}

func (t *TileMap) GetTile(x, y int) *Tile {
	return t.Tiles[x][y]
}

func (t *TileMap) Draw(screen *ebiten.Image) {
	for x := 0; x < t.Width/t.GridSize; x++ {
		for y := 0; y < t.Height/t.GridSize; y++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*t.GridSize), float64(y*t.GridSize))
			screen.DrawImage(t.Tiles[x][y].Image, op)

			if t.Tiles[x][y].Highlight {
				vector.DrawFilledRect(screen, float32(x*t.GridSize), float32(y*t.GridSize), float32(t.GridSize), float32(1), color.RGBA{255, 255, 255, 100}, true)
				vector.DrawFilledRect(screen, float32(x*t.GridSize), float32(y*t.GridSize), float32(1), float32(t.GridSize), color.RGBA{255, 255, 255, 100}, true)
				vector.DrawFilledRect(screen, float32(x*t.GridSize), float32((y+1)*t.GridSize-1), float32(t.GridSize), float32(1), color.RGBA{255, 255, 255, 100}, true)
				vector.DrawFilledRect(screen, float32((x+1)*t.GridSize-1), float32(y*t.GridSize), float32(1), float32(t.GridSize), color.RGBA{255, 255, 255, 100}, true)
			}
		}
	}
}

type Tile struct {
	Type      TileType
	X, Y      int
	Image     *ebiten.Image
	Walkable  bool
	Highlight bool
}

func NewTile(t TileType, x, y int, image *ebiten.Image, walkable bool) *Tile {
	return &Tile{
		Type:      t,
		X:         x,
		Y:         y,
		Image:     image,
		Walkable:  walkable,
		Highlight: false,
	}
}

type TileType int

const (
	GrassTile TileType = iota
	FarmTile
	PlantedTile
	WheatTile
)
