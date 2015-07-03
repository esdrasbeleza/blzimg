package operations

import (
	"image"
	"image/color"
	"image/draw"
	"testing"

	"github.com/esdrasbeleza/blzimg/containers"
	"github.com/stretchr/testify/assert"
)

func TestIfWeGetTheRightModeColorBetweenSomeOfThem(t *testing.T) {
	operation := ModeOperation{}

	pixel1 := color.RGBA{10, 100, 180, 0}
	pixel2 := color.RGBA{10, 120, 190, 0}
	pixel3 := color.RGBA{20, 120, 220, 0}
	pixel4 := color.RGBA{33, 130, 220, 0}

	result := color.RGBA{10, 120, 220, 0}

	assert.Equal(t, result, operation.mode([]color.Color{pixel1, pixel2, pixel3, pixel4}), "The mode color is (10, 120, 220)")
}

func TestIfWeGetAnImageMadeWithTheModePixelsIfWeMergeSomeImages(t *testing.T) {
	var (
		operation = ModeOperation{}

		black = color.RGBA{0, 0, 0, 0}
		white = color.RGBA{255, 255, 255, 0}
	)

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

	var (
		imageContainer1 = containers.ImageItselfContainer{image1}
		imageContainer2 = containers.ImageItselfContainer{image2}
		imageContainer3 = containers.ImageItselfContainer{image3}

		mergedImage, _ = operation.Result([]containers.ImageContainer{imageContainer1, imageContainer2, imageContainer3})
		bounds         = mergedImage.Bounds().Canon()
	)

	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			assert.Equal(t, black, mergedImage.At(i, j), "Pixel must be black!")
		}
	}
}
