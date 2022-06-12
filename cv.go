/* Computer Vision Model using Go. */
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

const (
	// The number of goroutines to use.
	// Try playing around with this number!
	goroutines = 4
)

func main() {
	// Get the filepath of the image.
	_, currentfile, _, _ := runtime.Caller(0)
	imagefile := filepath.Join(filepath.Dir(currentfile), "../../images/gopher.png")

	// Open the file.
	file, err := os.Open(imagefile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the image.
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// Get the bounds of the image.
	b := img.Bounds()

	// Make a new grayscale image of the same size.
	gray := image.NewGray(b)

	// Create a wait group so we can process the pixels in parallel.
	var wg sync.WaitGroup
	wg.Add(goroutines)

	// Calculate the stride (the number of pixels in each row).
	stride := b.Max.X / goroutines

	// Loop over the goroutines.
	for i := 0; i < goroutines; i++ {
		// Calculate the minimum and maximum X coordinates.
		minX := i * stride
		maxX := (i + 1) * stride

		// Create a new anonymous goroutine.
		go func() {
			// Loop over the Y coordinates.
			for y := b.Min.Y; y < b.Max.Y; y++ {
				// Loop over the X coordinates.
				for x := minX; x < maxX; x++ {
					// Get the color of the pixel.
					c := color.GrayModel.Convert(img.At(x, y))

					// Set the grayscale value.
					gray.Set(x, y, c)
				}
			}

			// Tell the wait group that we're done.
			wg.Done()
		}()
	}

	// Wait for the goroutines to finish.
	wg.Wait()

	// Encode the grayscale image and write it out to a new PNG file.
	outfile := filepath.Join(filepath.Dir(currentfile), "../../images/gopher-gray.png")
	out, err := os.Create(outfile)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	png.Encode(out, gray)

	// Print the path to the new image.
	fmt.Println("Created:", outfile)
}
