package operations

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"sort"

	"github.com/esdrasbeleza/blzimg/containers"
)

type Channel struct {
	Value uint8
	Count int
}

type ByFreq []Channel

func (c ByFreq) Len() int {
	return len(c)
}

func (c ByFreq) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (a ByFreq) Less(i, j int) bool {
	return a[i].Count < a[j].Count
}

type ModeOperation struct{}

func (m ModeOperation) mode(colors []color.Color) color.Color {
	var (
		rMap = make(map[uint32]int)
		gMap = make(map[uint32]int)
		bMap = make(map[uint32]int)
	)

	for _, currentColor := range colors {
		r, g, b, _ := currentColor.RGBA()
		rMap[r]++
		gMap[g]++
		bMap[b]++
	}

	var (
		rMean = modeForMap(rMap)
		gMean = modeForMap(gMap)
		bMean = modeForMap(bMap)
	)

	return color.RGBA{rMean, gMean, bMean, 0}
}

func modeForMap(cMap map[uint32]int) uint8 {
	size := len(cMap)
	count := make([]Channel, size)
	i := 0

	for k, v := range cMap {
		count[i] = Channel{
			Value: uint8(k),
			Count: v,
		}
		i++
	}

	sort.Sort(ByFreq(count))
	modeForMap := count[size-1].Value

	return modeForMap
}

func (c ModeOperation) Result(images []containers.ImageContainer) (image.Image, error) {
	imageCount := len(images)
	if imageCount == 0 {
		return nil, nil
	} else if imageCount == 1 {
		return images[0].GetImage(), nil
	}

	firstImage := images[0].GetImage()
	bounds := firstImage.Bounds()
	mode := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(mode, bounds, firstImage, bounds.Min, draw.Src)

	pixelTable := NewPixelDataTable(imageCount, bounds.Dx(), bounds.Dy())

	// Store data from all images -- Danger: O(n^3)
	imgCount := 0
	for _, currentImageContainer := range images {
		imgCount++
		fmt.Printf("\nReading image %d...\n", imgCount)

		currentImage := currentImageContainer.GetImage()
		if currentImage.Bounds() != bounds {
			return nil, errors.New("The images have different size!")
		}

		for i := currentImage.Bounds().Min.X; i < currentImage.Bounds().Max.X; i++ {
			for j := currentImage.Bounds().Min.Y; j < currentImage.Bounds().Max.Y; j++ {
				currentColor := currentImage.At(i, j)
				pixelTable.StoreData(i, j, currentColor)
			}
		}
	}

	fmt.Println("Generating output image...")
	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			currentPixelData := pixelTable.GetData(i, j)
			result := c.mode(currentPixelData.Colors)
			mode.Set(i, j, result)
		}
	}

	return mode, nil
}

type PixelData struct {
	Colors        []color.Color
	currentLength uint
}

type PixelDataTable struct {
	length    int
	dataTable map[image.Point]*PixelData
}

func NewPixelDataTable(imageCount, width, height int) *PixelDataTable {
	table := &PixelDataTable{length: imageCount}
	table.dataTable = make(map[image.Point]*PixelData)
	return table
}

func (ds *PixelDataTable) StoreData(x, y int, pointColor color.Color) error {
	point := image.Pt(x, y)
	pixelData := ds.dataTable[point]

	if pixelData == nil {
		pixelData = &PixelData{}
		pixelData.Colors = make([]color.Color, ds.length)
		pixelData.currentLength = 0

		ds.dataTable[point] = pixelData
	}

	pixelData.Colors[pixelData.currentLength] = pointColor
	pixelData.currentLength++

	return nil
}

func (ds *PixelDataTable) GetData(x, y int) *PixelData {
	point := image.Pt(x, y)
	return ds.dataTable[point]
}
