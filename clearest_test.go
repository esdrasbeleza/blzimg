package main

import (
	"github.com/stretchr/testify/assert"
	"image"
	"image/color"
	"image/draw"
	"testing"
)

func TestIfWeGetTheRightClearestColorBetweenSomeOfThem(t *testing.T) {
	operation := ClearestOperation{}

	pixel1 := color.RGBA{128, 22, 33, 0}
	pixel2 := color.RGBA{2, 21, 12, 0}
	pixel3 := color.RGBA{43, 12, 1, 0}
	pixel4 := color.RGBA{255, 255, 255, 0}

	assert.Equal(t, pixel4, operation.clearest([]color.RGBA{pixel1, pixel2, pixel3, pixel4}), "The clearest color is (128, 128, 128)")
}

func TestClearestWithEmptyArray(t *testing.T) {
	operation := ClearestOperation{}

	black := color.RGBA{0, 0, 0, 0}

	assert.Equal(t, black, operation.clearest([]color.RGBA{}), "Return black for an empty array")
}

func TestIfWeGetAImageMadeWithTheClearestPixelsIfWeMergeSomeImages(t *testing.T) {
	operation := ClearestOperation{}

	black := color.RGBA{0, 0, 0, 0}
	white := color.RGBA{255, 255, 255, 0}

	/*
	 * wbb
	 * wbb
	 * wbb
	 */
	image1 := image.NewRGBA(image.Rect(0, 0, 3, 3))
	draw.Draw(image1, image1.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)
	image1.Set(0, 0, white)
	image1.Set(0, 1, white)
	image1.Set(0, 2, white)

	/*
	 * bwb
	 * bwb
	 * bwb
	 */
	image2 := image.NewRGBA(image.Rect(0, 0, 3, 3))
	draw.Draw(image2, image2.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)
	image2.Set(1, 0, white)
	image2.Set(1, 1, white)
	image2.Set(1, 2, white)

	/*
	 * bbw
	 * bbw
	 * bbw
	 */
	image3 := image.NewRGBA(image.Rect(0, 0, 3, 3))
	draw.Draw(image3, image3.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)
	image3.Set(2, 0, white)
	image3.Set(2, 1, white)
	image3.Set(2, 2, white)

	mergedImage, _ := operation.Result([]image.Image{image1, image2, image3})
	bounds := mergedImage.Bounds().Canon()
	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			assert.Equal(t, white, mergedImage.At(i, j), "Pixel must be white!")
		}
	}
}
