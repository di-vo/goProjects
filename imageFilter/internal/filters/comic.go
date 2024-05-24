package internal

import (
    utils "assignments/imageFilter/internal/utils"
	"image"
	"image/color"
)

func ApplyComicFilter(img image.Image) image.Image {
	size := img.Bounds().Size()
	imgCopy := image.NewRGBA64(image.Rect(0, 0, size.X, size.Y))

	for x := range size.X {
		for y := range size.Y {
			r, g, b, a := img.At(x, y).RGBA()
            intens := utils.GetIntensity(r, g, b)

            switch true {
            case intens > 0 && intens <= 85:
                intens = 42
            case intens > 85 && intens <= 170:
                intens = 127
            case intens > 170 && intens <= 255:
                intens = 212
            }

			clr := color.RGBA64{
				R: uint16(intens),
				G: uint16(intens),
				B: uint16(intens),
				A: uint16(a),
			}
			imgCopy.SetRGBA64(x, y, clr)
		}
	}

	return imgCopy

}
