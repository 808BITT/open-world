package assets

import (
	"embed"
)

//go:embed player/standing/*.png
//go:embed player/walking/*.png
//go:embed tile/*.png
//go:embed wall/*.png
var Assets embed.FS

func EmbedAssets() embed.FS {
	return Assets
}
