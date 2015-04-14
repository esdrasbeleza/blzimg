package main

import "image"
import "image/draw"
import "image/color"
import "errors"
import "os"
import "fmt"

type ClearestOperation struct{}

func (c ClearestOperation) clearest(colors []color.RGBA) color.Color {
	clearest := color.RGBA{0, 0, 0, 0}

	for _, color := range colors {
		if average(color) > average(clearest) {
			clearest = color
		}
	}

	return clearest
}

func (c ClearestOperation) ResultFiles(files []string) (image.Image, error) {
	firstFile, _ := os.Open(files[0])
	defer firstFile.Close()
	firstImage, _, _ := image.Decode(firstFile)
	bounds := firstImage.Bounds()
	clearest := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(clearest, bounds, firstImage, bounds.Min, draw.Src)

	for _, currentImageFile := range files {
		fmt.Printf("Processing %s...\n", currentImageFile)

		currentFile, _ := os.Open(currentImageFile)
		defer currentFile.Close()
		currentImage, _, _ := image.Decode(currentFile)

		if currentImage.Bounds() != bounds {
			return nil, errors.New("The images have different size!")
		}

		imageToCompare := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
		draw.Draw(imageToCompare, bounds, currentImage, bounds.Min, draw.Src)
		c.getClearestImageBetweenTwo(clearest, imageToCompare)
	}

	return clearest, nil
}

func (c ClearestOperation) Result(images []image.Image) (image.Image, error) {
	firstImage := images[0]
	bounds := firstImage.Bounds()
	clearest := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(clearest, bounds, firstImage, bounds.Min, draw.Src)

	for _, currentImage := range images {
		if currentImage.Bounds() != bounds {
			return nil, errors.New("The images have different size!")
		}

		imageToCompare := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
		draw.Draw(imageToCompare, bounds, currentImage, bounds.Min, draw.Src)
		c.getClearestImageBetweenTwo(clearest, imageToCompare)
	}

	return clearest, nil
}

func (c ClearestOperation) getClearestImageBetweenTwo(current, other *image.RGBA) {
	for i := current.Bounds().Min.X; i < current.Bounds().Max.X; i++ {
		for j := current.Bounds().Min.Y; j < current.Bounds().Max.Y; j++ {
			currentClearestImagePixel := current.At(i, j).(color.RGBA)
			otherImagePixel := other.At(i, j).(color.RGBA)

			clearestColor := c.clearest([]color.RGBA{currentClearestImagePixel, otherImagePixel})
			current.Set(i, j, clearestColor)
		}
	}
}

func average(someColor color.Color) uint32 {
	r, g, b, _ := someColor.RGBA()
	return (r + g + b) / 3
}
