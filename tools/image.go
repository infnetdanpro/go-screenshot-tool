package tools

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/infnetdanpro/go-screenshot-tool/model"
	"github.com/kbinani/screenshot"
)

func WriteImage(maxWidth int, maxHeight int, bounds image.Rectangle) (*model.ImageObj, int, int, error) {
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return &model.ImageObj{}, maxWidth, maxHeight, err
	}

	newImg := image.NewRGBA(img.Bounds())
	draw.Draw(newImg, newImg.Bounds(), img, img.Bounds().Min, draw.Over)

	if err != nil {
		return &model.ImageObj{}, maxWidth, maxHeight, err
	}

	maxWidth, maxHeight = calculateSizes(bounds, maxWidth, maxHeight)
	data := &model.ImageObj{Image: img, Width: maxWidth, Height: maxHeight}

	return data, maxWidth, maxHeight, nil
}

func calculateSizes(bounds image.Rectangle, maxWidth int, maxHeight int) (int, int) {
	currentWidth := int(bounds.Dx())
	maxWidth = maxWidth + currentWidth

	currentHeight := int(bounds.Dy())
	if maxHeight < currentHeight {
		maxHeight = currentHeight
	}
	return maxWidth, maxHeight
}

func PrepareImage(maxWidth int, maxHeight int, images map[int]*model.ImageObj) (image.Rectangle, *image.RGBA) {
	maxPoint := image.Point{X: maxWidth, Y: maxHeight}
	minPoint := image.Point{X: 0, Y: 0}
	bounds := image.Rectangle{Min: minPoint, Max: maxPoint}

	wideImage := image.NewRGBA(bounds)
	draw.Draw(wideImage, wideImage.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src)

	globalMinXPoint := 0
	for _, imgData := range images {
		maxPoint := image.Point{X: imgData.Width + globalMinXPoint, Y: maxHeight}
		minPoint := image.Point{X: 0, Y: 0}
		bounds := image.Rectangle{Min: minPoint, Max: maxPoint}

		boundMinPoint := image.Point{X: -globalMinXPoint, Y: 0}
		draw.Draw(wideImage, bounds, imgData.Image, boundMinPoint, draw.Over)
		globalMinXPoint = globalMinXPoint + imgData.Width
	}
	return bounds, wideImage
}
