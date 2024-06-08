package internal

import (
	"assignments/imageFilter/internal/types"
	utils "assignments/imageFilter/internal/utils"
	"image"
	"image/color"
)

func applyBasicBlurFilter(img image.Image, imgCopy *image.RGBA64, d types.ImagePartData, x int, y int, radius int) {
	size := img.Bounds().Size()

	r, g, b, a := img.At(x, y).RGBA()

	colorsInRange := make([]color.Color, 0)

	// get every pixel in the range around the pixel
	for i := x - radius; i < (x-radius)+(radius*2+1); i++ {
		for j := y - radius; j < (y-radius)+(radius*2+1); j++ {
			if i-1 >= 0 && i+1 < size.X-1 && j-1 >= 0 && j+1 < size.Y-1 {
                // making sure to not include the pixel itself
                if i != x || j != y {
                    colorsInRange = append(colorsInRange, img.At(i, j))
                }
			}
		}
	}

	for _, e := range colorsInRange {
		r1, g1, b1, _ := e.RGBA()
		r += r1
		g += g1
		b += b1
	}

    // adding 1 for the pixel itself
	clr := color.RGBA64{
		R: uint16(int(r) / (len(colorsInRange) + 1)),
		G: uint16(int(g) / (len(colorsInRange) + 1)),
		B: uint16(int(b) / (len(colorsInRange) + 1)),
		A: uint16(a),
	}

	x = utils.MapToLocalCoords(x, d.Width, d.StartX)
	y = utils.MapToLocalCoords(y, d.Height, d.StartY)

	imgCopy.SetRGBA64(x, y, clr)
}
