package containers

import (
	"image"
)

type ImageItselfContainer struct {
	Image image.Image
}

func (i ImageItselfContainer) GetImage() image.Image {
	return i.Image
}
