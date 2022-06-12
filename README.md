# Computer Vision Model using Go

This code is a Go implementation of a computer vision model that converts an image to grayscale. The model is parallelized using goroutines, and the resulting grayscale image is written out to a new PNG file.

## Documentation

### `main` function

The `main` function is the entry point for the program. It starts by getting the filepath of the image to be processed. The image is then opened and decoded into an `image` object.

Next, the bounds of the image are calculated, and a new grayscale image is created. A wait group is also created so that the pixels can be processed in parallel.

Finally, the program loops over the goroutines, calculating the minimum and maximum X coordinates for each one. For each goroutine, a new anonymous goroutine is created. This goroutine loops over the Y coordinates, and for each Y coordinate it loops over the X coordinates. For each pixel, the color is converted to grayscale and set in the new grayscale image.

Once all of the pixels have been processed, the new grayscale image is encoded and written out to a new PNG file.

## Requirements

- Go

## Usage

To use this model, run the following command:

```
go run main.go
```

This will create a new image called `gopher-gray.png` in the `images` directory.
