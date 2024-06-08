package internal

import (
	"assignments/imageFilter/internal/types"
	utils "assignments/imageFilter/internal/utils"
	"image"
	"image/color"
	"math"
)

func boxBlurFilter(img image.Image, imgCopy *image.RGBA64, d types.ImagePartData, x int, y int, kernel int) {
	size := img.Bounds().Size()

	r, g, b, a := img.At(x, y).RGBA()

	colorsInRange := make([]color.Color, 0)

	// get every pixel in the range around the pixel
	for i := x - kernel; i < (x-kernel)+(kernel*2+1); i++ {
		for j := y - kernel; j < (y-kernel)+(kernel*2+1); j++ {
			if i-1 >= 0 && i+1 < size.X-1 && j-1 >= 0 && j+1 < size.Y-1 {
                // making sure to not include the pixel itself
                if i != x || j != y {
                    colorsInRange = append(colorsInRange, img.At(i, j))
                }
			}
		}
	}

	for _, e := range colorsInRange {
		r1, g1, b1, _ := e.RGBA()
		r += r1
		g += g1
		b += b1
	}

    // adding 1 for the pixel itself
	clr := color.RGBA64{
		R: uint16(int(r) / (len(colorsInRange) + 1)),
		G: uint16(int(g) / (len(colorsInRange) + 1)),
		B: uint16(int(b) / (len(colorsInRange) + 1)),
		A: uint16(a),
	}

	x = utils.MapToLocalCoords(x, d.Width, d.StartX)
	y = utils.MapToLocalCoords(y, d.Height, d.StartY)

	imgCopy.SetRGBA64(x, y, clr)
}

func gaussianBlur(img image.Image, imgCopy *image.RGBA64, d types.ImagePartData, x int, y int, kernel int) {
	size := img.Bounds().Size()

	r, g, b, a := img.At(x, y).RGBA()

	colorsInRange := make([]color.RGBA64, 0)
    contributionValue := 1.0

	// get every pixel in the range around the pixel
	for i := x - kernel; i < (x-kernel)+(kernel*2+1); i++ {
		for j := y - kernel; j < (y-kernel)+(kernel*2+1); j++ {
			if i-1 >= 0 && i+1 < size.X-1 && j-1 >= 0 && j+1 < size.Y-1 {
                // making sure to not include the pixel itself
                if i != x || j != y {
                    deltaW := math.Abs(float64(x - i))
                    deltaH := math.Abs(float64(y - j))

                    distance := math.Sqrt(math.Pow(deltaW, 2) + math.Pow(deltaH, 2))

                    r1, g1, b1, a1 := img.At(i, j).RGBA()

                    tclr := color.RGBA64{
                        R: uint16(float64(r1) / distance),
                        G: uint16(float64(g1) / distance),
                        B: uint16(float64(b1) / distance),
                        A: uint16(a1),
                    }

                    colorsInRange = append(colorsInRange, tclr)
                    contributionValue += 1 / distance
                }
			}
		}
	}

	for _, e := range colorsInRange {
		r1, g1, b1, _ := e.RGBA()
		r += r1
		g += g1
		b += b1
	}

    // adding 1 for the pixel itself
	clr := color.RGBA64{
		R: uint16(float64(r) / contributionValue),
		G: uint16(float64(g) / contributionValue),
		B: uint16(float64(b) / contributionValue),
		A: uint16(a),
	}

	x = utils.MapToLocalCoords(x, d.Width, d.StartX)
	y = utils.MapToLocalCoords(y, d.Height, d.StartY)

	imgCopy.SetRGBA64(x, y, clr)
}
