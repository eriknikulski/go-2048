package config

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
)

var (
	Cfg = pixelgl.WindowConfig{
		Title:  "2048",
		Bounds: pixel.R(0, 0, WindowSize.X, WindowSize.Y),
		VSync:  true,
	}
	WindowSize   = pixel.V(1024, 640)
	MainRectSize = pixel.V(503, 503)
	GridPadding  = 15.0
	TileRectVec  = pixel.V(107, 107)
	FontName     = "Roboto/Roboto-Bold.ttf"
	Background   = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	BigRect      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
)
