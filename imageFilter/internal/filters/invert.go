package internal

import (
	"image"
	"image/color"
)

func ApplyInvertFilter(img image.Image) image.Image {
    size := img.Bounds().Size()
	imgCopy := image.NewRGBA64(image.Rect(0, 0, size.X, size.Y))

	for x := range size.X {
		for y := range size.Y {
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
			imgCopy.SetRGBA64(x, y, clr)
		}
	}

    return imgCopy
}
