package internal

import (
	types "assignments/imageFilter/internal/types"
	utils "assignments/imageFilter/internal/utils"
	"image"
	"image/color"
)

func ApplyComicFilter(img image.Image, d types.ImagePartData, imageChan chan types.ImageData) {
	imgCopy := image.NewRGBA64(image.Rect(0, 0, d.Width, d.Height))

	for x := d.StartX; x < d.StartX + d.Width; x++ {
		for y := d.StartY; y < d.StartY + d.Height; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			intens := utils.GetIntensity(r, g, b) / 255

			switch true {
			case intens > 0 && intens <= 85:
				intens = 42
			case intens > 85 && intens <= 170:
				intens = 127
			case intens > 170 && intens <= 255:
				intens = 212 
			}

			clr := color.RGBA64{
				R: uint16(intens) * 255,
				G: uint16(intens) * 255,
				B: uint16(intens) * 255,
				A: uint16(a),
			}
            newX := x
            newY := y

            if newX >= d.Width {
                newX -= d.StartX
            }

            if newY >= d.Height {
                newY -= d.StartY
            }
			imgCopy.SetRGBA64(newX, newY, clr)
		}
	}

    out := types.ImageData{
        Img: imgCopy,
        StartX: d.StartX,
        StartY: d.StartY,
    }

	imageChan <- out
}
