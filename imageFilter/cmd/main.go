package main

import (
	filters "github.com/di-vo/goProjects/imageFilter/internal/filters"
	types "github.com/di-vo/goProjects/imageFilter/internal/types"
	utils "github.com/di-vo/goProjects/imageFilter/internal/utils"
	"flag"
	"fmt"
	"image"
	"runtime"
	"strings"
	"sync"
	"time"

	ansi "github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

var (
	filterFlag  = flag.String("f", "", "The filter to apply to the image. Valid filters are:\n boxBlur\n gaussianBlur\n edge\n spot\n invert\n comic\n heat\n sort (VERY SLOW!)\n pixel (experimental)\n basicKuwahara\n generalKuwahara (experimental)")
	sourceFlag  = flag.String("s", "", "The name of the source image file")
	helpFlag    = flag.Bool("h", false, "Shows this help message")
	convertFlag = flag.Bool("c", false, "Create a new PNG from the given JPEG image")
	threadsFlag = flag.Int("t", runtime.GOMAXPROCS(0), "Specify the number of threads used for the filter(0 < x <= image height)\nDefaults to the number of available threads")
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

	fmt.Printf("%s read in\n", format)

	var wg sync.WaitGroup

	imgWidth := img.Bounds().Size().X
	imgHeight := img.Bounds().Size().Y

	imgCopy := image.NewRGBA64(image.Rect(0, 0, imgWidth, imgHeight))

	partWidth := imgWidth
	partHeight := imgHeight / *threadsFlag

	imageChan := make(chan types.ImageData, 100)
	defer close(imageChan)

    filterStart := time.Now()

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

	bar := progressbar.NewOptions(int(*threadsFlag),
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionSetWidth(50),
		progressbar.OptionSetElapsedTime(true),
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionSetDescription("Applying filter..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "—",
			SaucerHead:    ">",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
        progressbar.OptionOnCompletion(func() { fmt.Printf("\nTime to apply the filter: %.4fs\n", time.Since(filterStart).Seconds()) }))

	go func() {
		for m := range imageChan {
			// fmt.Printf("image received: %d, %d\n", m.StartX, m.StartY)
			bar.Add(1)
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
	fmt.Printf("Done!\nTotal execution time: %.4fs\n", time.Since(start).Seconds())
}
