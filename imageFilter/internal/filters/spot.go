package internal

import (
	types "github.com/di-vo/goProjects/imageFilter/internal/types"
	utils "github.com/di-vo/goProjects/imageFilter/internal/utils"
	"image"
	"image/color"
	"math"
)

func getNewColor(c uint32, radius uint32, distance float64) uint32 {
	if uint32(distance) <= radius {
		return c - c/radius*uint32(distance)
	} else {
		return 0
	}
}

func spotFilter(img image.Image, imgCopy *image.RGBA64, d types.ImagePartData, x int, y int, radius uint32) {
	size := img.Bounds().Size()

	centerX := size.X / 2
	centerY := size.Y / 2

	r, g, b, a := img.At(x, y).RGBA()

	deltaW := math.Abs(float64(centerX - x))
	deltaH := math.Abs(float64(centerY - y))

	distance := math.Sqrt(math.Pow(deltaW, 2) + math.Pow(deltaH, 2))

	r = getNewColor(r, radius, distance)
	g = getNewColor(g, radius, distance)
	b = getNewColor(b, radius, distance)
	// a = a - a / radius * uint32(distance)

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
