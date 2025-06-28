package main

import (
	"fmt"
	"os"

	gopixels "go-pixels/internal/go_pixels"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	// define lipgloss style for output
	style := lipgloss.NewStyle().
		Padding(1)

	// renderer with custom width, height, brightness & style
	renderer := gopixels.NewPixelRenderer(80, 24, 128, style)

	imagePath := "/Users/saran13raj/Desktop/workspace/saran13raj/go-pixels/image.png"

	if len(imagePath) < 2 {
		fmt.Println("gopixels: Invalid image path")
		os.Exit(1)
	}

	// Example 1: Halfcell rendering with color
	fmt.Println("=== Halfcell Rendering (Color) ===")
	output, err := renderer.FromImagePath(imagePath, map[string]string{
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
	output, err = renderer.FromImagePath(imagePath, map[string]string{
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
	output, err = renderer.FromImagePath(imagePath, map[string]string{
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
	output, err = renderer.FromImagePath(imagePath, map[string]string{
		"type":  "fullcell",
		"color": "false",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "gopixels error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(output)
}
