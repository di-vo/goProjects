package internal

import (
	"assignments/imageFilter/internal/types"
	utils "assignments/imageFilter/internal/utils"
	"image"
	"image/color"
)

func invertFilter(img image.Image, imgCopy *image.RGBA64, d types.ImagePartData, x int, y int) {
	r, g, b, a := img.At(x, y).RGBA()

	r = (255 - (r / 255)) * 255
	g = (255 - (g / 255)) * 255
	b = (255 - (b / 255)) * 255

	clr := color.RGBA64{
		R: uint16(r),
		G: uint16(g),
		B: uint16(b),
		A: uint16(a),
	}

    x = utils.MapToLocalCoords(x, d.Width, d.StartX)
    y = utils.MapToLocalCoords(y, d.Height, d.StartY)

	imgCopy.SetRGBA64(x, y, clr)
}
