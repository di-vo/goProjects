package internal

import (
	"image"
	"image/color"
)

func getIntensity(r, g, b uint32) float64 {
	return 0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)
}

func ApplyEdgeFilter(img image.Image) image.Image {
	size := img.Bounds().Size()
	imgCopy := image.NewRGBA64(image.Rect(0, 0, size.X, size.Y))

	for x := range size.X {
		for y := range size.Y {
			r, g, b, a := img.At(x, y).RGBA()
            intens := getIntensity(r, g, b)
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
                intensN := getIntensity(r1, g1, b1) * 2
                intensS := getIntensity(r2, g2, b2) * 2
                intensW := getIntensity(r3, g3, b3) * 2
                intensE := getIntensity(r4, g4, b4) * 2

                intensNE := getIntensity(r5, g5, b5)
                intensSE := getIntensity(r6, g6, b6)
                intensSW := getIntensity(r7, g7, b7)
                intensNW := getIntensity(r8, g8, b8)

				r = uint32(intens) + (uint32(intensNW + intensW + intensSW) - uint32(intensNE + intensE + intensSE)) + (uint32(intensNW + intensN + intensNE) - uint32(intensSW + intensS + intensSE))
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
			imgCopy.SetRGBA64(x, y, clr)
		}
	}

	return imgCopy
}