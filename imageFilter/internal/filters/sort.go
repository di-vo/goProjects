package internal

import (
	"github.com/di-vo/goProjects/imageFilter/internal/types"
	utils "github.com/di-vo/goProjects/imageFilter/internal/utils"
	"image"
	"image/color"
	"sort"
)

func sortFilter(img image.Image, imgCopy *image.RGBA64, d types.ImagePartData, x int, y int) {
	// NOTE: New Idea (I know this is very slow but I can't think of a smarter solution rn):
	// 1. Make array containing colors of the row
	// 2. Sort array based on intensity
	// 3. Use x to get correct color

	rowColors := make([]color.RGBA64, img.Bounds().Size().X)

	for i := range img.Bounds().Size().X {
		r, g, b, a := img.At(i, y).RGBA()

		clr := color.RGBA64{
			R: uint16(r),
			G: uint16(g),
			B: uint16(b),
			A: uint16(a),
		}

		rowColors[i] = clr
	}

	sort.Slice(rowColors, func(i, j int) bool {
		r1, g1, b1, _ := rowColors[i].RGBA()
		r2, g2, b2, _ := rowColors[j].RGBA()

        intens1 := utils.GetIntensity(r1, b1, g1)
        intens2 := utils.GetIntensity(r2, b2, g2)

        // hue1 := utils.GetHue(r1, g1, b1)
        // hue2 := utils.GetHue(r2, g2, b2)

		return intens1 < intens2
	})

    // when I tried accessing it directly from the slice, I would always get an out of bounds error,
    // so I'm using a little workaround here
    rowColorsMap := make(map[int]color.RGBA64)

    for i, e := range rowColors {
        rowColorsMap[i] = e

        if i > x {
            break
        }
    }

    clr := rowColorsMap[x]

	tx := utils.MapToLocalCoords(x, d.Width, d.StartX)
	ty := utils.MapToLocalCoords(y, d.Height, d.StartY)

	imgCopy.SetRGBA64(tx, ty, clr)
}
