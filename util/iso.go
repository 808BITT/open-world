package util

func GridToIso(x, y, tileWidth, tileHeight, screenWidth int) (int, int) {
	return (screenWidth-tileWidth)/2 + (x-y)*(tileWidth/2), (x+y)*(tileHeight/4) + 200
}

func IsoToGrid(x, y, tileWidth, tileHeight, screenWidth int) (int, int) {
	if x >= (screenWidth-tileWidth)/2 {
		x -= (screenWidth-tileWidth)/2 + x/tileWidth
	} else {
		x -= (screenWidth-tileWidth)/2 + tileWidth
	}
	y -= 200

	return ((x/(tileWidth/2) + y/(tileHeight/4)) / 2), ((y/(tileHeight/4) - (x / (tileWidth / 2))) / 2)
}
