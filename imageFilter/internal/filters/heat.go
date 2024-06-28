package internal

import (
	"github.com/di-vo/goProjects/imageFilter/internal/types"
	utils "github.com/di-vo/goProjects/imageFilter/internal/utils"
	"image"
	"image/color"
)

func heatFilter(img image.Image, imgCopy *image.RGBA64, d types.ImagePartData, x int, y int) {
	r, g, b, a := img.At(x, y).RGBA()
	intens := utils.GetIntensity(r, g, b) / 255

	switch true {
	case intens > 0 && intens <= 42:
		r = 0
		g = 0
		b = 0
	case intens > 42 && intens <= 84:
		r = 0
		g = 0
		b = 255
	case intens > 84 && intens <= 126:
		r = 0
		g = 255
		b = 255
	case intens > 126 && intens <= 168:
		r = 0
		g = 255
		b = 0
	case intens > 168 && intens <= 210:
		r = 255
		g = 255
		b = 0
	case intens > 210 && intens <= 255:
		r = 255
		g = 0
		b = 0
	}

	clr := color.RGBA64{
		R: uint16(r * 255),
		G: uint16(g * 255),
		B: uint16(b * 255),
		A: uint16(a),
	}

    x = utils.MapToLocalCoords(x, d.Width, d.StartX)
    y = utils.MapToLocalCoords(y, d.Height, d.StartY)

	imgCopy.SetRGBA64(x, y, clr)
}
