package internal

import (
    utils "assignments/imageFilter/internal/utils"
	"image"
	"image/color"
)

func ApplyHeatFilter(img image.Image) image.Image {
	size := img.Bounds().Size()
	imgCopy := image.NewRGBA64(image.Rect(0, 0, size.X, size.Y))

	for x := range size.X {
		for y := range size.Y {
			r, g, b, a := img.At(x, y).RGBA()
            intens := utils.GetIntensity(r, g, b) / 255

            switch true {
            case intens > 0 && intens <= 42:
                r = 0
                g = 0
                b = 0
            case intens > 42 && intens <= 84:
                r = 0
                g = 0
                b = 255
            case intens > 84 && intens <= 126:
                r = 0
                g = 255
                b = 255
            case intens > 126 && intens <= 168:
                r = 0
                g = 255
                b = 0
            case intens > 168 && intens <= 210:
                r = 255
                g = 255
                b = 0
            case intens > 210 && intens <= 255:
                r = 255
                g = 0
                b = 0
            }

			clr := color.RGBA64{
                R: uint16(r * 255),
                G: uint16(g * 255),
                B: uint16(b * 255),
				A: uint16(a),
			}
			imgCopy.SetRGBA64(x, y, clr)
		}
	}

	return imgCopy

}
