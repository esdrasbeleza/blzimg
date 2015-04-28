package main

import (
	"image"
	"os"
)

type ImageContainer interface {
	getImage() image.Image
}

type FileImageContainer struct {
	filename string
}

func (f FileImageContainer) getImage() image.Image {
	file, _ := os.Open(f.filename)
	defer file.Close()
	image, _, _ := image.Decode(file)
	return image
}

type ImageItselfContainer struct {
	image image.Image
}

func (i ImageItselfContainer) getImage() image.Image {
	return i.image
}
