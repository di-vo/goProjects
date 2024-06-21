package internal

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func DecodeFile(fileName string) (image.Image, string, error) {
	fi, err := os.Open(fileName)
	if err != nil {
		return nil, "", err
	}
	defer fi.Close()

	img, format, err := image.Decode(fi)
	if err != nil {
		return nil, "", err
	}

	return img, format, nil
}

func SaveNewImage(img image.Image, outputName string) {
	outFile, err := os.Create(outputName)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	png.Encode(outFile, img)
}

// Helper function to generate a usable png image from a jpeg image
func JpegToPngConv(fileName string) {
	fi, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	img, err := jpeg.Decode(fi)
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create(strings.Split(fileName, ".")[0] + ".png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	png.Encode(outFile, img)
}

func GetIntensity(r, g, b uint32) float64 {
	return (0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b))
}

func GetHue(r, g, b uint32) float64 {
	clrMin := min(r, g, b)
	clrMax := max(r, g, b)

    if clrMin == clrMax {
        return 0
    }

    var hue float64

    if clrMax == r {
        hue = float64((g - b) / (clrMax - clrMin))
    } else if clrMax == g {
        hue = float64(2 + (b - r) / (clrMax - clrMin))
    } else {
        hue = float64(4 + (r - g) / (clrMax - clrMin))
    }

    hue *= 60

    if hue < 0 {
        hue += 360
    }

    return hue
}

func MapToLocalCoords(val int, length int, offset int) int {
	if val >= length {
		return val - offset
	}

	return val
}
