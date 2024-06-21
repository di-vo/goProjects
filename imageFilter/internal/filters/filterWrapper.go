package internal

import (
	types "assignments/imageFilter/internal/types"
	"image"
)

func ApplyFilter(filter string, img image.Image, d types.ImagePartData, imageChan chan types.ImageData) {
	imgCopy := image.NewRGBA64(image.Rect(0, 0, d.Width, d.Height))

	for x := d.StartX; x < d.StartX+d.Width; x++ {
		for y := d.StartY; y < d.StartY+d.Height; y++ {
			switch filter {
			case "comic":
				comicFilter(img, imgCopy, d, x, y)
			case "spot":
				spotFilter(img, imgCopy, d, x, y, 700)
			case "invert":
				invertFilter(img, imgCopy, d, x, y)
			case "edge":
				edgeFilter(img, imgCopy, d, x, y)
			case "boxBlur":
				boxBlurFilter(img, imgCopy, d, x, y, 4)
			case "gaussianBlur":
				gaussianBlur(img, imgCopy, d, x, y, 4)
			case "heat":
				heatFilter(img, imgCopy, d, x, y)
			case "sort":
				sortFilter(img, imgCopy, d, x, y)
			case "pixel":
				pixelFilter(img, imgCopy, d, x, y, 2)
			case "basicKuwahara":
				basicKuwahara(img, imgCopy, d, x, y, 4)
            case "generalKuwahara":
                generalKuwahara(img, imgCopy, d, x, y, 4)
			}
		}
	}

	out := types.ImageData{
		Img:    imgCopy,
		StartX: d.StartX,
		StartY: d.StartY,
	}

	imageChan <- out
}
