package internal

import (
	types "assignments/imageFilter/internal/types"
	utils "assignments/imageFilter/internal/utils"
	"image"
	"image/color"
)

func applyComicFilter(img image.Image, imgCopy *image.RGBA64, d types.ImagePartData, x int, y int) {
	r, g, b, a := img.At(x, y).RGBA()
	intens := utils.GetIntensity(r, g, b) / 255

	switch true {
	case intens > 0 && intens <= 85:
		intens = 42
	case intens > 85 && intens <= 170:
		intens = 127
	case intens > 170 && intens <= 255:
		intens = 212
	}

	clr := color.RGBA64{
		R: uint16(intens) * 255,
		G: uint16(intens) * 255,
		B: uint16(intens) * 255,
		A: uint16(a),
	}

    x = utils.MapToLocalCoords(x, d.Width, d.StartX)
    y = utils.MapToLocalCoords(y, d.Height, d.StartY)

	imgCopy.SetRGBA64(x, y, clr)
}
