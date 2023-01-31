package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/infnetdanpro/go-screenshot-tool/model"
	"github.com/infnetdanpro/go-screenshot-tool/tools"
	"github.com/kbinani/screenshot"
)

var quality int
var display int
var delay int
var customPath string

func init() {
	flag.IntVar(&quality, "quality", 70, "0/100 quality of image. Example: -quality=70")
	flag.IntVar(&display, "display", 0, "number of display, 1,2,3 etc or 0 (for all). Example: -display=2")
	flag.IntVar(&delay, "delay", 0, "delay screenshot in seconds, 1,2,3 etc. Example: -delay=5")
	flag.StringVar(&customPath, "path", "", "path to save screenshot. Example: -path=/tmp/")
	flag.Parse()
}

func main() {
	displays := detectDisplay()

	images := make(map[int]*model.ImageObj)

	// detect max
	maxWidth := 0
	maxHeight := 0

	if delay > 0 {
		tools.Delay(delay)
	}

	for i := 0; i < displays; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		imgData, newMaxWidth, newMaxHeight, err := tools.WriteImage(maxWidth, maxHeight, bounds)
		if err != nil {
			fmt.Println("Can't process image")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		maxWidth = newMaxWidth
		maxHeight = newMaxHeight

		images[i] = imgData
	}

	// Prepare white background full image
	// Write each screenshot to the white image
	bounds, wideImage := tools.PrepareImage(maxWidth, maxHeight, images)

	filename, err := tools.SaveImage(customPath, maxWidth, maxHeight, wideImage, bounds, quality)
	if err != nil {
		fmt.Println("Can't save the file!")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(filename)
}

func detectDisplay() int {
	var displays int
	if display < 1 {
		displays = screenshot.NumActiveDisplays()
	} else {
		displays = display
	}
	return displays
}
