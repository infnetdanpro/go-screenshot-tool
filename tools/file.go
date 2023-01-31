package tools

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"time"
)

func getCurrentDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return currentDir, nil
}

func SaveImage(customPath string, maxWidth int, maxHeight int, wideImage *image.RGBA, bounds image.Rectangle, quality int) (string, error) {
	pathToSave := customPath

	if len(customPath) == 0 {
		currentPath, err := getCurrentDir()
		if err != nil {
			return "", err
		}
		pathToSave = currentPath
	}

	timestamp := time.Now().UnixNano()

	fileName := fmt.Sprintf("Screenshot_%d_%dx%d.jpg", timestamp, maxWidth, maxHeight)

	path := filepath.Join(pathToSave, fileName)
	file, err := os.Create(path)
	defer file.Close()

	if err != nil {
		return "", err
	}

	var opt jpeg.Options
	opt.Quality = quality

	err = jpeg.Encode(file, wideImage, &opt)
	if err != nil {
		return "", err
	}
	return path, nil
}
