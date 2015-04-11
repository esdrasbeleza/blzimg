package blzimg

import "image"
import "image/draw"
import "image/color"
import "errors"

type ClearestOperation struct{}

func (c ClearestOperation) Clearest(colors []color.RGBA) color.Color {
	clearest := color.RGBA{0, 0, 0, 0}

	for _, color := range colors {
		if average(color) > average(clearest) {
			clearest = color
		}
	}

	return clearest
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

			clearestColor := c.Clearest([]color.RGBA{currentClearestImagePixel, otherImagePixel})
			current.Set(i, j, clearestColor)
		}
	}
}

func average(color color.Color) uint32 {
	r, g, b, _ := color.RGBA()
	return (r + g + b) / 3
}
