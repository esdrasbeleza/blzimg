package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"image/jpeg"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "blzimg"
	app.Usage = "Execute some operations on images"
	app.Version = "0.1"

	app.Commands = []cli.Command{
		{
			Name:    "lightest",
			Aliases: []string{"l"},
			Usage:   "Merge the lightest pixels of some images in a single one",
			Action: func(c *cli.Context) {
				fmt.Printf("Processing images...")
				operation := ClearestOperation{}
				finalImage, _ := operation.ResultFiles(c.Args())
				finalFile, _ := os.Create("final.jpg")
				defer finalFile.Close()
				jpeg.Encode(finalFile, finalImage, &jpeg.Options{jpeg.DefaultQuality})
				fmt.Printf(" done.\n")
			},
		},
	}

	app.Run(os.Args)
}
