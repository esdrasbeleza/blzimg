package blzimg

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"testing"
)

func TestClearest(t *testing.T) {
	operation := ClearestOperation{}

	pixel1 := color.RGBA{128, 22, 33, 0}
	pixel2 := color.RGBA{2, 21, 12, 0}
	pixel3 := color.RGBA{43, 12, 1, 0}
	pixel4 := color.RGBA{255, 255, 255, 0}

	assert.Equal(t, pixel4, operation.Clearest([]color.RGBA{pixel1, pixel2, pixel3, pixel4}), "The clearest color is (128, 128, 128)")
}

func TestClearestWithEmptyArray(t *testing.T) {
	operation := ClearestOperation{}

	black := color.RGBA{0, 0, 0, 0}

	assert.Equal(t, black, operation.Clearest([]color.RGBA{}), "Return black for an empty array")
}
