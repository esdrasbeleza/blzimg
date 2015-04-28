# blzimg

A command line utility written in Go to do some image operations.

## Why not Photoshop, The GIMP or ImageMagick?

I know that these tools do what I want to do here and much more. But I wanted to understand some
algorithms and learn some Go.

## Using

Up to this moment, the only operation in blzimg is the *lightest* operation. This operation reads
every pixel from a list of images and, for a position (x,y), the final image will have the lightest
pixel at that position.

The command is

`blzimg --output final.jpg lightest image1.jpg image2.jpg image3.jpg`

### More information about *lightest*

Please read [this post](http://wp.me/pMrQd-7H).

## Limitations

Until now, blzimg is working only with images in JPEG format.

