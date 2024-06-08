package types

import (
	"image"
)

type ImageData struct {
	Img    image.Image
	StartX int
	StartY int
}

type ImagePartData struct {
    StartX int
    StartY int
    Width int
    Height int
}

type Vec2 struct {
    X int
    Y int
}
