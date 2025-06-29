package main

import (
	"fmt"
	"image"
	"os"

	gopixels "github.com/saran13raj/go-pixels"

	"golang.org/x/image/webp"
)

func main() {
	imagePath := "tmp/image.png"

	img, err := loadWebP("tmp/parrot.webp")

	// Example 1: Halfcell rendering with color from webp stream
	fmt.Println("=== Halfcell Rendering (Color) ===")
	output, err := gopixels.FromImageStream(img, 50, 55, "halfcell", true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("\n" + output)

	// Example 2: Fullcell rendering with color
	fmt.Println("\n=== Fullcell Rendering (Color) ===")
	output, err = gopixels.FromImagePath(imagePath, 0, 0, "fullcell", true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("\n" + output)

	// Example 3: Halfcell rendering with greyscale
	fmt.Println("=== Halfcell Rendering (Greyscale) ===")
	output, err = gopixels.FromImagePath(imagePath, 0, 0, "halfcell", false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("\n" + output)

	// Example 4: Fullcell rendering with greyscale from image stream
	fmt.Println("\n=== Fullcell Rendering (Greyscale) ===")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading avif: %v\n", err)
		os.Exit(1)
	}
	output, err = gopixels.FromImageStream(img, 70, 75, "fullcell", false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("\n" + output)
}

func loadWebP(path string) (image.Image, error) {
	// Open the WebP file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the WebP file into an image.Image
	img, err := webp.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}
