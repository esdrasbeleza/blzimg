package blzimg

import "image/color"

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

func average(color color.Color) uint32 {
	r, g, b, _ := color.RGBA()
	return (r + g + b) / 3
}
