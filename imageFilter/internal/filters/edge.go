package internal

import (
	"assignments/imageFilter/internal/types"
	utils "assignments/imageFilter/internal/utils"
	"image"
	"image/color"
)

func applyEdgeFilter(img image.Image, imgCopy *image.RGBA64, d types.ImagePartData, x int, y int) {
	size := img.Bounds().Size()

	r, g, b, a := img.At(x, y).RGBA()
	intens := utils.GetIntensity(r, g, b)
	var count uint32 = 9

	if y-1 >= 0 && y+1 < size.Y-1 && x-1 >= 0 && x+1 < size.X-1 {
		r1, g1, b1, _ := img.At(x, y-1).RGBA()
		r2, g2, b2, _ := img.At(x, y+1).RGBA()
		r3, g3, b3, _ := img.At(x-1, y).RGBA()
		r4, g4, b4, _ := img.At(x+1, y).RGBA()

		r5, g5, b5, _ := img.At(x+1, y-1).RGBA()
		r6, g6, b6, _ := img.At(x+1, y+1).RGBA()
		r7, g7, b7, _ := img.At(x-1, y+1).RGBA()
		r8, g8, b8, _ := img.At(x-1, y-1).RGBA()

		// weighing direct neighbors twice as much
		intensN := utils.GetIntensity(r1, g1, b1) * 2
		intensS := utils.GetIntensity(r2, g2, b2) * 2
		intensW := utils.GetIntensity(r3, g3, b3) * 2
		intensE := utils.GetIntensity(r4, g4, b4) * 2

		intensNE := utils.GetIntensity(r5, g5, b5)
		intensSE := utils.GetIntensity(r6, g6, b6)
		intensSW := utils.GetIntensity(r7, g7, b7)
		intensNW := utils.GetIntensity(r8, g8, b8)

		r = uint32(intens) + (uint32(intensNW+intensW+intensSW) - uint32(intensNE+intensE+intensSE)) + (uint32(intensNW+intensN+intensNE) - uint32(intensSW+intensS+intensSE))
		g = r
		b = r
	} else {
		r = uint32(intens)
		g = r
		b = r
	}

	clr := color.RGBA64{
		R: uint16(r / count),
		G: uint16(g / count),
		B: uint16(b / count),
		A: uint16(a),
	}

    x = utils.MapToLocalCoords(x, d.Width, d.StartX)
    y = utils.MapToLocalCoords(y, d.Height, d.StartY)

	imgCopy.SetRGBA64(x, y, clr)
}
