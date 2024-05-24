package main

import (
	filters "assignments/imageFilter/internal/filters"
	utils "assignments/imageFilter/internal/utils"
	"flag"
	"fmt"
	"image"
	"strings"
)


var (
    filterFlag = flag.String("f", "", "Define a filter. Valid filters are blur, edge, spot, invert, comic.")
    imageFlag = flag.String("i", "", "The name of the image file")
    helpFlag = flag.Bool("h", false, "Shows this help message")
    outputName string
    filterName string
)

func main() {
	flag.Parse()

    if *helpFlag {
        flag.PrintDefaults()
        return
    }

	// utils.JpegToPngConv("mountain.jpg")

	img, format, err := utils.DecodeFile(*imageFlag)
	if err != nil {
		panic(err)
	}
	fmt.Printf("format: %s\n", format)

    var imgCopy image.Image

    switch *filterFlag {
    case "spot":
	    imgCopy = filters.ApplySpotFilter(img, 2000)
    case "blur":
	    imgCopy = filters.ApplyBasicBlurFilter(img, 20)
    case "edge":
	    imgCopy = filters.ApplyEdgeFilter(img)
    case "invert":
	    imgCopy = filters.ApplyInvertFilter(img)
    case "comic":
	    imgCopy = filters.ApplyComicFilter(img)
    }

	utils.SaveNewImage(imgCopy, fmt.Sprintf("%s_%s.png", strings.Split(*imageFlag, ".")[0], *filterFlag))
}
