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
				applyComicFilter(img, imgCopy, d, x, y)
			case "spot":
				applySpotFilter(img, imgCopy, d, x, y, 700)
			case "invert":
				applyInvertFilter(img, imgCopy, d, x, y)
			case "edge":
				applyEdgeFilter(img, imgCopy, d, x, y)
			case "blur":
                applyBasicBlurFilter(img, imgCopy, d, x, y, 4)
            case "heat":
                applyHeatFilter(img, imgCopy, d, x, y)
            case "sort":
                applySortFilter(img, imgCopy, d, x, y)
            case "pixel":
                applyPixelFilter(img, imgCopy, d, x, y, 2)
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
