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
