package assets

import (
	"embed"
)

type Assets struct {
	embed.FS
}

//go:embed tile/*.png
//go:embed wall/*.png
//go:embed player/standing/*.png
//go:embed player/walking/*.png
//go:embed test/*.png
var files embed.FS

func EmbedAssets() *Assets {
	return &Assets{files}
}
