package operations

import (
	"errors"
	"image"
	"image/color"
	"image/draw"

	"github.com/esdrasbeleza/blzimg/containers"
)

type LightestOperation struct{}

func (c LightestOperation) lightest(color1, color2 color.Color) color.Color {
	if c.luminance(color1) > c.luminance(color2) {
		return color1
	} else {
		return color2
	}
}

func (c LightestOperation) Result(images []containers.ImageContainer) (image.Image, error) {
	if len(images) == 0 {
		return nil, nil
	} else if len(images) == 1 {
		return images[0].GetImage(), nil
	}

	firstImage := images[0].GetImage()
	bounds := firstImage.Bounds()
	lightest := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(lightest, bounds, firstImage, bounds.Min, draw.Src)

	for _, currentImageContainer := range images[1:] {
		currentImage := currentImageContainer.GetImage()
		if currentImage.Bounds() != bounds {
			return nil, errors.New("The images have different size!")
		}

		c.getLightestImageBetweenTwo(lightest, currentImage)
	}

	return lightest, nil
}

func (c LightestOperation) getLightestImageBetweenTwo(current *image.RGBA, other image.Image) {
	for i := current.Bounds().Min.X; i < current.Bounds().Max.X; i++ {
		for j := current.Bounds().Min.Y; j < current.Bounds().Max.Y; j++ {
			currentLightestImagePixel := current.At(i, j)
			otherImagePixel := other.At(i, j)

			lightestColor := c.lightest(currentLightestImagePixel, otherImagePixel)
			if currentLightestImagePixel != lightestColor {
				current.Set(i, j, lightestColor)
			}
		}
	}
}

// http://stackoverflow.com/questions/596216/formula-to-determine-brightness-of-rgb-color
func (c LightestOperation) luminance(someColor color.Color) uint32 {
	r, g, b, _ := someColor.RGBA()
	return uint32(0.2126*float32(r) + 0.7152*float32(g) + 0.0722*float32(b))
}
