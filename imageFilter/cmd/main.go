package main

import (
	"fmt"
    filters "assignments/imageFilter/internal/filters"
    utils "assignments/imageFilter/internal/utils"
)

const (
	fileName   = "bridge.png"
    outputName = "bridge_spot.png"
)


func main() {
    // utils.JpegToPngConv("bridge.jpeg")

	img, format, err := utils.DecodeFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("format: %s\n", format)

    imgCopy := filters.ApplySpotFilter(img, 2000)

    utils.SaveNewImage(imgCopy, outputName)
}
