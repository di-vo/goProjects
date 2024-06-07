package main

import (
	filters "assignments/imageFilter/internal/filters"
	types "assignments/imageFilter/internal/types"
	utils "assignments/imageFilter/internal/utils"
	"flag"
	"fmt"
	"image"
	"strings"
	"sync"
)

var (
    filterFlag = flag.String("f", "", "Define a filter. Valid filters are blur, edge, spot, invert, comic, heat.")
    sourceFlag = flag.String("s", "", "The name of the source image file")
    helpFlag = flag.Bool("h", false, "Shows this help message")
    convertFlag = flag.Bool("c", false, "Create a new PNG from the given JPEG image")
    outputName string
    filterName string
)

func main() {
	flag.Parse()

    if *helpFlag {
        flag.PrintDefaults()
        return
    }

    if *convertFlag {
	    utils.JpegToPngConv(*sourceFlag)
        return
    }

	img, format, err := utils.DecodeFile(*sourceFlag)
	if err != nil {
		panic(err)
	}
	fmt.Printf("format: %s\n", format)

    imageChan := make(chan types.ImageData)
    var wg sync.WaitGroup
    // defer close(imageChan)

    // 100 x 100

    // 0, 0, 50, 50
    // 50, 0, 50, 50
    // 0, 50, 50, 50
    // 50, 50, 50, 50
    imgWidth := img.Bounds().Size().X
    imgHeight := img.Bounds().Size().Y

	imgCopy := image.NewRGBA64(image.Rect(0, 0, imgWidth, imgHeight))

    partWidth := imgWidth / 2
    partHeight := imgHeight / 2

    for x := 0; x < imgWidth; x += partWidth {
        for y := 0; y < imgHeight; y += partHeight {
            data := types.ImagePartData{
                StartX: x,
                StartY: y,
                Width: partWidth,
                Height: partHeight,
            }

            wg.Add(1)

            switch *filterFlag {
            case "spot":
                go func() {
                    filters.ApplySpotFilter(img, data, imageChan, 600)
                    wg.Done()
                }()
            case "comic":
                go func() {
                    filters.ApplyComicFilter(img, data, imageChan)
                    wg.Done()
                }()
            // case "blur":
            //  imgCopy = filters.ApplyBasicBlurFilter(img, 10)
            // case "edge":
            //  imgCopy = filters.ApplyEdgeFilter(img)
            // case "invert":
            //  imgCopy = filters.ApplyInvertFilter(img)
            // case "heat":
            //  imgCopy = filters.ApplyHeatFilter(img)
            }
        }
    }

    go func() {
        for m := range imageChan {
            fmt.Printf("image received: %d, %d\n", m.StartX, m.StartY)
            for x := range m.Img.Bounds().Size().X {
                for y := range m.Img.Bounds().Size().Y {
                    imgCopy.Set(m.StartX + x, m.StartY + y, m.Img.At(x, y))
                }
            }
        }
    }()

    wg.Wait()

	utils.SaveNewImage(imgCopy, fmt.Sprintf("%s_%s.png", strings.Split(*sourceFlag, ".")[0], *filterFlag))
}
