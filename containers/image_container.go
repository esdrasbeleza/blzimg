package containers

import (
	"image"
)

type ImageContainer interface {
	GetImage() image.Image
}
