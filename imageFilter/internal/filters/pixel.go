package internal

import (
	"assignments/imageFilter/internal/types"
	utils "assignments/imageFilter/internal/utils"
	"image"
	"image/color"
)

func applyPixelFilter(img image.Image, imgCopy *image.RGBA64, d types.ImagePartData, x int, y int, radius int) {
	r, g, b, a := img.At(x, y).RGBA()

	clr := color.RGBA64{
		R: uint16(r),
		G: uint16(g),
		B: uint16(b),
		A: uint16(a),
	}

	if x > 0 && y > 0 && x%(radius*2+1) == 0 && y%(radius*2+1) == 0 {
		size := img.Bounds().Size()

		for i := x - radius; i < (x-radius)+(radius*2+1); i++ {
			for j := y - radius; j < (y-radius)+(radius*2+1); j++ {
				if i-1 >= 0 && i+1 < size.X-1 && j-1 >= 0 && j+1 < size.Y-1 {
					tx := utils.MapToLocalCoords(i, d.Width, d.StartX)
					ty := utils.MapToLocalCoords(j, d.Height, d.StartY)

					imgCopy.SetRGBA64(tx, ty, clr)
				}
			}
		}
	} else {
		x = utils.MapToLocalCoords(x, d.Width, d.StartX)
		y = utils.MapToLocalCoords(y, d.Height, d.StartY)

		imgCopy.SetRGBA64(x, y, clr)
	}

}
