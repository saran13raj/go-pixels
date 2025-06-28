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
	output, err := gopixels.FromImagePath(imagePath, 50, 55, "halfcell", true)
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

	// Example 1: Halfcell rendering with greyscale
	fmt.Println("=== Halfcell Rendering (Greyscale) ===")
	output, err = gopixels.FromImagePath(imagePath, 0, 0, "halfcell", false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("\n" + output)

	// Example 2: Fullcell rendering with greyscale
	fmt.Println("\n=== Fullcell Rendering (Greyscale) ===")
	output, err = gopixels.FromImagePath(imagePath, 40, 45, "fullcell", false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("\n" + output)
}
