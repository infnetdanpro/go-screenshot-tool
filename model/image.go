package model

import "image"

type ImageObj struct {
	Image  *image.RGBA
	Width  int
	Height int
}
