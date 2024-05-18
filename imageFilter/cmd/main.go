package main

import (
	"fmt"
    filters "assignments/imageFilter/internal/filters"
    utils "assignments/imageFilter/internal/utils"
)

const (
	fileName   = "papagei.png"
    outputName = "papagei_edge.png"
)


func main() {
    // utils.JpegToPngConv("bridge.jpeg")

	img, format, err := utils.DecodeFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("format: %s\n", format)

    imgCopy := filters.ApplyEdgeFilter(img)

    utils.SaveNewImage(imgCopy, outputName)
}
