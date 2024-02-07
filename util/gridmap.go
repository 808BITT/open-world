package util

import (
	"open-world/assets"
)

type Point struct {
	X, Y int
}

type IsoPoint struct {
	P Point
}

type GridPoint struct {
	P Point
}

type GridMap struct {
	Lookup *map[IsoPoint]GridPoint
}

func NewGridMap(tileShape []Point, gW, gH, tW, screenW int, a *assets.Assets) *GridMap {
	gridMap := make(map[IsoPoint]GridPoint)
	for x := 0; x < gW; x++ {
		for y := 0; y < gH; y++ {
			grid := GridPoint{P: Point{x, y}}
			iX, iY := GridToIso(x, y, 2, tW, screenW)
			for _, p := range tileShape {
				iso := IsoPoint{P: Point{iX + p.X, iY + p.Y}}
				if _, ok := gridMap[iso]; !ok {
					gridMap[iso] = grid
				}
			}
		}
	}
	return &GridMap{
		Lookup: &gridMap,
	}
}
