package main

import (
	"fmt"
	"image/jpeg"
	"os"

	"github.com/codegangsta/cli"
	"github.com/esdrasbeleza/blzimg/containers"
	"github.com/esdrasbeleza/blzimg/operations"
)

func main() {
	app := cli.NewApp()
	app.Name = "blzimg"
	app.Usage = "Execute some operations on images"
	app.Version = "0.1"

	app.Authors = []cli.Author{cli.Author{"Esdras Beleza", "esdras@esdrasbeleza.com"}}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output",
			Value: "final.jpg",
			Usage: "Output file",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "lightest",
			Aliases: []string{"l"},
			Usage:   "Merge the lightest pixels of some images in a single one",
			Action: func(c *cli.Context) {
				filenames := c.Args()
				output := c.GlobalString("output")

				fmt.Printf("Processing images...")
				fileContainers := make([]containers.ImageContainer, len(filenames))
				for index, filename := range filenames {
					fileContainers[index] = containers.FileImageContainer{filename}
				}

				operation := operations.LightestOperation{}
				finalImage, _ := operation.Result(fileContainers)
				finalFile, _ := os.Create(output)
				defer finalFile.Close()
				jpeg.Encode(finalFile, finalImage, &jpeg.Options{jpeg.DefaultQuality})
				fmt.Printf(" done.\n")
			},
		},
		{
			Name:    "mode",
			Aliases: []string{"m"},
			Usage:   "Merge the lightest pixels of some images in a single one",
			Action: func(c *cli.Context) {
				filenames := c.Args()
				output := c.GlobalString("output")

				fmt.Printf("Processing images...")
				fileContainers := make([]containers.ImageContainer, len(filenames))
				for index, filename := range filenames {
					fileContainers[index] = containers.FileImageContainer{filename}
				}

				operation := operations.ModeOperation{}
				finalImage, _ := operation.Result(fileContainers)
				finalFile, _ := os.Create(output)
				defer finalFile.Close()
				jpeg.Encode(finalFile, finalImage, &jpeg.Options{jpeg.DefaultQuality})
				fmt.Printf(" done.\n")
			},
		},
	}

	app.Run(os.Args)
}
