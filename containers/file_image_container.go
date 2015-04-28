package containers

import (
	"image"
	"os"
)

type FileImageContainer struct {
	Filename string
}

func (f FileImageContainer) GetImage() image.Image {
	file, _ := os.Open(f.Filename)
	defer file.Close()
	image, _, _ := image.Decode(file)
	return image
}
