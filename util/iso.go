package util

func GridToIso(x, y, z, tileSize, screenWidth int) (int, int) {
	return (screenWidth-tileSize)/2 + (x-y)*(tileSize/2), (x+y)*(tileSize/4) + 150 - z*tileSize/2
}
