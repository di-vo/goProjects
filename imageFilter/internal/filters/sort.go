package internal

import (
	"assignments/imageFilter/internal/types"
	utils "assignments/imageFilter/internal/utils"
	"image"
	"image/color"
)

func sortFilter(img image.Image, imgCopy *image.RGBA64, d types.ImagePartData, x int, y int) {
	r, g, b, a := img.At(x, y).RGBA()

	intens := utils.GetIntensity(r, g, b)

	for i := range img.Bounds().Size().X {
		r1, g1, b1, a1 := img.At(i, y).RGBA()

		intens1 := utils.GetIntensity(r1, g1, b1)

		if intens1 > intens {
			clr := color.RGBA64{
				R: uint16(r),
				G: uint16(g),
				B: uint16(b),
				A: uint16(a),
			}

            tx := utils.MapToLocalCoords(i, d.Width, d.StartX)
            ty := utils.MapToLocalCoords(y, d.Height, d.StartY)

			imgCopy.SetRGBA64(tx, ty, clr)

			clr = color.RGBA64{
				R: uint16(r1),
				G: uint16(g1),
				B: uint16(b1),
				A: uint16(a1),
			}

			tx = utils.MapToLocalCoords(x, d.Width, d.StartX)
			ty = utils.MapToLocalCoords(y, d.Height, d.StartY)

			imgCopy.SetRGBA64(tx, ty, clr)
            break
		}
	}
}
