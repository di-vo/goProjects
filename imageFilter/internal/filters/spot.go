package internal

import (
	"image"
	"image/color"
	"math"
)

func getNewColor(c uint32, radius uint32, distance float64) uint32 {
    if uint32(distance) <= radius {
        return c - c / radius * uint32(distance)
    } else {
        return 0
    }
}

func ApplySpotFilter(img image.Image, radius uint32) image.Image {
    size := img.Bounds().Size()
	imgCopy := image.NewRGBA64(image.Rect(0, 0, size.X, size.Y))

    centerX := size.X / 2
    centerY := size.Y / 2

	for x := range size.X {
		for y := range size.Y {
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
			imgCopy.SetRGBA64(x, y, clr)
		}
	}

    return imgCopy
}
