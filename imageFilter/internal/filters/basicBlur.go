package internal

import (
	"image"
	"image/color"
)

func ApplyBasicBlurFilter(img image.Image, iterations int) image.Image {
	size := img.Bounds().Size()
	imgCopy := image.NewRGBA64(image.Rect(0, 0, size.X, size.Y))

	for range iterations {
		for x := range size.X {
			for y := range size.Y {
				r, g, b, a := img.At(x, y).RGBA()
				var count uint32 = 1

				if y-1 >= 0 {
					r1, g1, b1, _ := img.At(x, y-1).RGBA()
					r += r1
					g += g1
					b += b1
					count++
				}

				if x+1 < size.X-1 {
					r1, g1, b1, _ := img.At(x+1, y).RGBA()
					r += r1
					g += g1
					b += b1
					count++
				}

				if y+1 < size.Y-1 {
					r1, g1, b1, _ := img.At(x, y+1).RGBA()
					r += r1
					g += g1
					b += b1
					count++
				}

				if x-1 >= 0 {
					r1, g1, b1, _ := img.At(x-1, y).RGBA()
					r += r1
					g += g1
					b += b1
					count++
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
        img = imgCopy
	}

	return imgCopy
}
