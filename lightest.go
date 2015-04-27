package main

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
)

type LightestOperation struct{}

func (c LightestOperation) lightest(color1, color2 color.RGBA) color.Color {
	if c.luminance(color1) > c.luminance(color2) {
		return color1
	} else {
		return color2
	}
}

func (c LightestOperation) Result(images []ImageContainer) (image.Image, error) {
	if len(images) == 0 {
		return nil, nil
	} else if len(images) == 1 {
		return images[0].getImage(), nil
	} else if len(images) == 2 {
		firstImage := images[0].getImage()
		bounds := firstImage.Bounds()
		lightest := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
		draw.Draw(lightest, bounds, firstImage, bounds.Min, draw.Src)

		secondImage := images[1].getImage()
		if secondImage.Bounds() != bounds {
			return nil, errors.New("The images have different size!")
		}

		imageToCompare := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
		draw.Draw(imageToCompare, bounds, secondImage, bounds.Min, draw.Src)

		c.getLightestImageBetweenTwo(lightest, imageToCompare)

		return lightest, nil

	} else {
		type tuple struct {
			image image.Image
			err   error
		}

		imageChan := make(chan tuple)
		defer close(imageChan)

		go func() {
			img, err := c.Result(images[:len(images)/2])
			imageChan <- tuple{img, err}
		}()

		go func() {
			img, err := c.Result(images[len(images)/2:])
			imageChan <- tuple{img, err}
		}()

		tuple1 := <-imageChan
		if tuple1.err != nil {
			return nil, tuple1.err
		}

		tuple2 := <-imageChan
		if tuple2.err != nil {
			return nil, tuple2.err
		}

		lightest1 := ImageItselfContainer{tuple1.image}
		lightest2 := ImageItselfContainer{tuple2.image}

		return c.Result([]ImageContainer{lightest1, lightest2})
	}

	return nil, nil
}

func (c LightestOperation) getLightestImageBetweenTwo(lightestImage, other *image.RGBA) {
	for i := lightestImage.Bounds().Min.X; i < lightestImage.Bounds().Max.X; i++ {
		for j := lightestImage.Bounds().Min.Y; j < lightestImage.Bounds().Max.Y; j++ {
			pixelFromLightestImage := lightestImage.At(i, j).(color.RGBA)
			pixelFromOtherImage := other.At(i, j).(color.RGBA)

			lightestColor := c.lightest(pixelFromLightestImage, pixelFromOtherImage)
			if pixelFromLightestImage != lightestColor {
				lightestImage.Set(i, j, lightestColor)
			}
		}
	}
}

// http://stackoverflow.com/questions/596216/formula-to-determine-brightness-of-rgb-color
func (c LightestOperation) luminance(someColor color.Color) uint32 {
	r, g, b, _ := someColor.RGBA()
	return uint32(0.2126*float32(r) + 0.7152*float32(g) + 0.0722*float32(b))
}

func (c LightestOperation) average(someColor color.Color) uint32 {
	r, g, b, _ := someColor.RGBA()
	return (r + g + b) / 3
}
