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
	"time"
)

var (
    filterFlag  = flag.String("f", "", "Define a filter. Valid filters are:\n boxBlur\n gaussianBlur\n edge\n spot\n invert\n comic\n heat\n sort(experimental)\n pixel(experimental)")
	sourceFlag  = flag.String("s", "", "The name of the source image file")
	helpFlag    = flag.Bool("h", false, "Shows this help message")
	convertFlag = flag.Bool("c", false, "Create a new PNG from the given JPEG image")
	outputName  string
	filterName  string
)

func main() {
    start := time.Now()
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
	defer close(imageChan)

	var wg sync.WaitGroup

	// 100 x 100

	// 0, 0, 50, 50
	// 50, 0, 50, 50
	// 0, 50, 50, 50
	// 50, 50, 50, 50
	imgWidth := img.Bounds().Size().X
	imgHeight := img.Bounds().Size().Y

	imgCopy := image.NewRGBA64(image.Rect(0, 0, imgWidth, imgHeight))

	partWidth := imgWidth / 10
	partHeight := imgHeight / 10

	for x := 0; x < imgWidth; x += partWidth {
		for y := 0; y < imgHeight; y += partHeight {
			data := types.ImagePartData{
				StartX: x,
				StartY: y,
				Width:  partWidth,
				Height: partHeight,
			}

			wg.Add(1)
			go filters.ApplyFilter(*filterFlag, img, data, imageChan)
		}
	}

	go func() {
		for m := range imageChan {
			// fmt.Printf("image received: %d, %d\n", m.StartX, m.StartY)
			for x := range m.Img.Bounds().Size().X {
				for y := range m.Img.Bounds().Size().Y {
					imgCopy.Set(m.StartX+x, m.StartY+y, m.Img.At(x, y))
				}
			}
			wg.Done()
		}
	}()

	wg.Wait()

	utils.SaveNewImage(imgCopy, fmt.Sprintf("%s_%s.png", strings.Split(*sourceFlag, ".")[0], *filterFlag))
    end := time.Since(start)
    fmt.Printf("Execution Time: %.4fs\n", end.Seconds())
}
