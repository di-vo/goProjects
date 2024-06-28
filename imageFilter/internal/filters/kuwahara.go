package internal

import (
	"github.com/di-vo/goProjects/imageFilter/internal/types"
	utils "github.com/di-vo/goProjects/imageFilter/internal/utils"
	"image"
	"image/color"
	"math"
)

func setQuadrantDataBox(img image.Image, startX int, startY int, centerX int, centerY int, kernel int, quadrants map[float64]color.RGBA64) {
	_, _, _, a := img.At(centerX, centerY).RGBA()
	size := img.Bounds().Size()
	var sumR, sumG, sumB uint32
	var deviation float64

	colorsInRange := make([]color.Color, 0)
	colorsInRange = append(colorsInRange, img.At(centerX, centerY))

	for i := startX; i < startX+(kernel+1); i++ {
		for j := startY; j < startY+(kernel+1); j++ {
			if i-1 >= 0 && i+1 < size.X-1 && j-1 >= 0 && j+1 < size.Y-1 {
				if i != centerX || j != centerY {
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
}

type sector struct {
	colors            []color.Color
	contributionValue float64
	deviation         float64
	sumR, sumG, sumB, sumA  uint32
	averageColor      color.Color
}

func getSector(centerX int, centerY int, x int, y int) int {
	theta := math.Atan2(float64(y-centerY), float64(x-centerX))

	if theta < 0 {
		theta += 2 * math.Pi
	}

	theta -= math.Pi / 8

	if theta < 0 {
		theta += 2 * math.Pi
	}

	sectorAngle := 2 * math.Pi / 8

	sector := int(math.Floor(theta / sectorAngle))

	return sector
}

func getGeneralizedColor(img image.Image, startX int, startY int, centerX int, centerY int, kernel int) color.RGBA64 {
	_, _, _, a := img.At(centerX, centerY).RGBA()
	size := img.Bounds().Size()

	var sectors [8]sector

    // initialize sectors
	for _, s := range sectors {
		s.colors = append(s.colors, img.At(centerX, centerY))
		s.contributionValue = 1.0
	}

	for i := startX; i < startX+(kernel*2+1); i++ {
		for j := startY; j < startY+(kernel*2+1); j++ {
			if i-1 >= 0 && i+1 < size.X-1 && j-1 >= 0 && j+1 < size.Y-1 {
				deltaW := math.Abs(float64(centerX - i))
				deltaH := math.Abs(float64(centerY - j))

				distance := math.Sqrt(math.Pow(deltaW, 2) + math.Pow(deltaH, 2))
				if (i != centerX || j != centerY) && distance <= float64(kernel+1) {
					r1, g1, b1, a1 := img.At(i, j).RGBA()

					tclr := color.RGBA64{
						R: uint16(float64(r1) / distance),
						G: uint16(float64(g1) / distance),
						B: uint16(float64(b1) / distance),
						A: uint16(float64(a1) / distance),
					}

					s := getSector(centerX, centerY, i, j)
					sectors[s].colors = append(sectors[s].colors, tclr)
					sectors[s].contributionValue += 1 / distance
				}
			}
		}
	}

    // calculate new color
	sumWeights := 0.0
	var sumR, sumG, sumB, sumA uint32

	for _, s := range sectors {
		for _, e := range s.colors {
			r1, g1, b1, a1 := e.RGBA()
			s.sumR += r1
			s.sumG += g1
			s.sumB += b1
            s.sumA += a1

			intensity := utils.GetIntensity(r1, g1, b1)
			s.deviation += intensity
		}

		s.deviation /= float64(len(s.colors))

		s.averageColor = color.RGBA64{
			R: uint16(float64(s.sumR) / s.contributionValue),
			G: uint16(float64(s.sumG) / s.contributionValue),
			B: uint16(float64(s.sumB) / s.contributionValue),
			A: uint16(float64(s.sumA) / s.contributionValue),
		}

		sumWeights += 1/(1 + s.deviation)
		sumR += uint32((float64(s.sumR) / s.contributionValue) * (1/(1 + s.deviation)))
		sumG += uint32((float64(s.sumG) / s.contributionValue) * (1/(1 + s.deviation)))
		sumB += uint32((float64(s.sumB) / s.contributionValue) * (1/(1 + s.deviation)))
		sumA += uint32((float64(s.sumA) / s.contributionValue) * (1/(1 + s.deviation)))
	}

	return color.RGBA64{
		R: uint16(float64(sumR) / sumWeights),
		G: uint16(float64(sumG) / sumWeights),
		B: uint16(float64(sumB) / sumWeights),
		A: uint16(a),
	}
}

func basicKuwaharaFilter(img image.Image, imgCopy *image.RGBA64, d types.ImagePartData, x int, y int, kernel int) {
    // NOTE: explanation:
    // 1. define a sqare kernel around the pixel
    // 2. split the kernel into four quadrants
    // 3. calculate the average color and deviation for each quadrant
    // 4. find the quadrant with the lowest deviation, and use its color for the pixel

	quadrants := make(map[float64]color.RGBA64)

	setQuadrantDataBox(img, x-kernel, y-kernel, x, y, kernel, quadrants)
	setQuadrantDataBox(img, x, y-kernel, x, y, kernel, quadrants)
	setQuadrantDataBox(img, x-kernel, y, x, y, kernel, quadrants)
	setQuadrantDataBox(img, x, y, x, y, kernel, quadrants)

	//finding lowest deviation
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

func generalKuwaharaFilter(img image.Image, imgCopy * image.RGBA64, d types.ImagePartData, x int, y int, kernel int) {
    // NOTE: explanation:
    // 1. define a circular kernel around the pixel
    // 2. divide the kernel into eight sectors
    // 3. calculate average color for each sector, taking into account the distance to the center pixel
    // 4. use a specific formula to calculate the new color to apply to the pixel

	clr := getGeneralizedColor(img, x-kernel, y-kernel, x, y, kernel)

	x = utils.MapToLocalCoords(x, d.Width, d.StartX)
	y = utils.MapToLocalCoords(y, d.Height, d.StartY)

	imgCopy.SetRGBA64(x, y, clr)
}
