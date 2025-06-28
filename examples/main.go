package main

import (
	"fmt"
	"os"

	gopixels "github.com/saran13raj/go-pixels"
)

func main() {

	imagePath := "tmp/image.png"

	if len(imagePath) < 2 {
		fmt.Println("gopixels: Invalid image path")
		os.Exit(1)
	}

	// Example 1: Halfcell rendering with color
	fmt.Println("=== Halfcell Rendering (Color) ===")
	output, err := gopixels.FromImagePath(imagePath, 50, 55, map[string]string{
		"type":  "halfcell",
		"color": "true",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "gopixels error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(output)

	// Example 2: Fullcell rendering with color
	fmt.Println("\n=== Fullcell Rendering (Color) ===")
	output, err = gopixels.FromImagePath(imagePath, 0, 0, map[string]string{
		"type":  "fullcell",
		"color": "true",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "gopixels error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(output)

	// Example 1: Halfcell rendering with greyscale
	fmt.Println("=== Halfcell Rendering (Greyscale) ===")
	output, err = gopixels.FromImagePath(imagePath, 0, 0, map[string]string{
		"type":  "halfcell",
		"color": "false",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "gopixels error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(output)

	// Example 2: Fullcell rendering with greyscale
	fmt.Println("\n=== Fullcell Rendering (Greyscale) ===")
	output, err = gopixels.FromImagePath(imagePath, 40, 45, map[string]string{
		"type":  "fullcell",
		"color": "false",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "gopixels error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(output)
}
