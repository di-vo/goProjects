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

    var imgCopy image.Image

    switch *filterFlag {
    case "spot":
	    imgCopy = filters.ApplySpotFilter(img, 200)
    case "blur":
	    imgCopy = filters.ApplyBasicBlurFilter(img, 20)
    case "edge":
	    imgCopy = filters.ApplyEdgeFilter(img)
    case "invert":
	    imgCopy = filters.ApplyInvertFilter(img)
    case "comic":
	    imgCopy = filters.ApplyComicFilter(img)
    case "heat":
	    imgCopy = filters.ApplyHeatFilter(img)
    }

	utils.SaveNewImage(imgCopy, fmt.Sprintf("%s_%s.png", strings.Split(*sourceFlag, ".")[0], *filterFlag))
}
