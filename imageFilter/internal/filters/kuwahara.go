package internal

import (
	"assignments/imageFilter/internal/types"
	utils "assignments/imageFilter/internal/utils"
	"image"
	"image/color"
)

func basicKuwahara(img image.Image, imgCopy *image.RGBA64, d types.ImagePartData, x int, y int, kernel int) {
	size := img.Bounds().Size()

	_, _, _, a := img.At(x, y).RGBA()

	quadrants := make(map[float64]color.RGBA64)
	var sumR, sumG, sumB uint32
	var deviation float64

	colorsInRange := make([]color.Color, 0)
	colorsInRange = append(colorsInRange, img.At(x, y))

	// NW quad
	for i := x - kernel; i < (x-kernel)+(kernel+1); i++ {
		for j := y - kernel; j < (y-kernel)+(kernel+1); j++ {
			if i-1 >= 0 && i+1 < size.X-1 && j-1 >= 0 && j+1 < size.Y-1 {
				if i != x || j != y {
					colorsInRange = append(colorsInRange, img.At(i, j))
				}
			}
		}
	}

	for _, e := range colorsInRange {
		r1, g1, b1, _ := e.RGBA()
		sumR += r1
		sumG += g1
		sumB += b1

		intensity := utils.GetIntensity(r1, g1, b1)
		deviation += intensity
	}

	quadrants[deviation/float64(len(colorsInRange))] = color.RGBA64{
		R: uint16(int(sumR) / len(colorsInRange)),
		G: uint16(int(sumG) / len(colorsInRange)),
		B: uint16(int(sumB) / len(colorsInRange)),
		A: uint16(a),
	}

	sumR, sumG, sumB = 0, 0, 0
	deviation = 0.0
	colorsInRange = make([]color.Color, 0)
	colorsInRange = append(colorsInRange, img.At(x, y))

	// NE quad
	for i := x; i < x+kernel+1; i++ {
		for j := y - kernel; j < (y-kernel)+(kernel+1); j++ {
			if i-1 >= 0 && i+1 < size.X-1 && j-1 >= 0 && j+1 < size.Y-1 {
				if i != x || j != y {
					colorsInRange = append(colorsInRange, img.At(i, j))
				}
			}
		}
	}

	for _, e := range colorsInRange {
		r1, g1, b1, _ := e.RGBA()
		sumR += r1
		sumG += g1
		sumB += b1

		intensity := utils.GetIntensity(r1, g1, b1)
		deviation += intensity
	}

	quadrants[deviation/float64(len(colorsInRange))] = color.RGBA64{
		R: uint16(int(sumR) / len(colorsInRange)),
		G: uint16(int(sumG) / len(colorsInRange)),
		B: uint16(int(sumB) / len(colorsInRange)),
		A: uint16(a),
	}

	sumR, sumG, sumB = 0, 0, 0
	deviation = 0.0
	colorsInRange = make([]color.Color, 0)
	colorsInRange = append(colorsInRange, img.At(x, y))

	// SW quad
	for i := x - kernel; i < (x-kernel)+(kernel+1); i++ {
		for j := y; j < y+kernel+1; j++ {
			if i-1 >= 0 && i+1 < size.X-1 && j-1 >= 0 && j+1 < size.Y-1 {
				if i != x || j != y {
					colorsInRange = append(colorsInRange, img.At(i, j))
				}
			}
		}
	}

	for _, e := range colorsInRange {
		r1, g1, b1, _ := e.RGBA()
		sumR += r1
		sumG += g1
		sumB += b1

		intensity := utils.GetIntensity(r1, g1, b1)
		deviation += intensity
	}

	quadrants[deviation/float64(len(colorsInRange))] = color.RGBA64{
		R: uint16(int(sumR) / len(colorsInRange)),
		G: uint16(int(sumG) / len(colorsInRange)),
		B: uint16(int(sumB) / len(colorsInRange)),
		A: uint16(a),
	}

	sumR, sumG, sumB = 0, 0, 0
	deviation = 0.0
	colorsInRange = make([]color.Color, 0)
	colorsInRange = append(colorsInRange, img.At(x, y))

	// SE quad
	for i := x; i < x+kernel+1; i++ {
		for j := y; j < y+kernel+1; j++ {
			if i-1 >= 0 && i+1 < size.X-1 && j-1 >= 0 && j+1 < size.Y-1 {
				if i != x || j != y {
					colorsInRange = append(colorsInRange, img.At(i, j))
				}
			}
		}
	}

	for _, e := range colorsInRange {
		r1, g1, b1, _ := e.RGBA()
		sumR += r1
		sumG += g1
		sumB += b1

		intensity := utils.GetIntensity(r1, g1, b1)
		deviation += intensity
	}

	quadrants[deviation/float64(len(colorsInRange))] = color.RGBA64{
		R: uint16(int(sumR) / len(colorsInRange)),
		G: uint16(int(sumG) / len(colorsInRange)),
		B: uint16(int(sumB) / len(colorsInRange)),
		A: uint16(a),
	}

	// finding lowest deviation
	devs := make([]float64, 0)
	for i := range quadrants {
		devs = append(devs, i)
	}

	smallestDev := devs[0]

	for _, e := range devs {
		if e < smallestDev {
			smallestDev = e
		}
	}

	x = utils.MapToLocalCoords(x, d.Width, d.StartX)
	y = utils.MapToLocalCoords(y, d.Height, d.StartY)

	imgCopy.SetRGBA64(x, y, quadrants[smallestDev])
}
