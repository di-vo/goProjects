package main

import (
	"fmt"
    filters "assignments/imageFilter/internal/filters"
    utils "assignments/imageFilter/internal/utils"
)

const (
	fileName   = "papagei.png"
	outputName = "papagei_blur.png"
)


func main() {
	img, format, err := utils.DecodeFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("format: %s\n", format)

    imgCopy := filters.ApplyBasicBlurFilter(img, 10)

    utils.SaveNewImage(imgCopy, outputName)
}
