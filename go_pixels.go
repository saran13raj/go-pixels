package gopixels

import (
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/saran13raj/go-pixels/renderer"

	"golang.org/x/image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// loads and decodes an image from a file path.
func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return img, nil
}

// scales the image to the specified width and height.
func resizeImage(img image.Image, width, height int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.NearestNeighbor.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)
	return dst
}

// FromImagePath converts an image to string representation
func FromImagePath(path string, width, height int, renderType string, useColor bool) (string, error) {
	img, err := loadImage(path)
	if err != nil {
		return "", fmt.Errorf("gopixels failed to load content (%v)", err)
	}

	// Get image dimensions for aspect ratio calculation
	bounds := img.Bounds()
	imgWidth := bounds.Dx()
	imgHeight := bounds.Dy()
	aspectRatio := float64(imgWidth) / float64(imgHeight)

	// Set defaults based on terminal size and image aspect ratio
	terminalWidth := 80  // default terminal width
	terminalHeight := 24 // default terminal height

	if width <= 0 || height <= 0 {
		// Calculate aspect ratio with respect to terminal dimensions
		if width <= 0 && height <= 0 {
			// Both unspecified, use terminal size maintaining aspect ratio
			targetWidth := terminalWidth
			targetHeight := int(float64(targetWidth) / aspectRatio)
			if targetHeight > terminalHeight {
				targetHeight = terminalHeight
				targetWidth = int(float64(targetHeight) * aspectRatio)
			}
			width = targetWidth
			height = targetHeight
		} else if width <= 0 {
			// Width unspecified, calculate from height maintaining aspect ratio
			width = int(float64(height) * aspectRatio)
		} else if height <= 0 {
			// Height unspecified, calculate from width maintaining aspect ratio
			height = int(float64(width) / aspectRatio)
		}
	}

	if renderType == "" {
		renderType = "halfcell"
	}
	defaultColor := "" // TODO: pass in color to render

	var output string
	var resized image.Image

	switch renderType {
	case "halfcell":
		// For halfcell, use the specified height directly without adjustment
		resized = resizeImage(img, width, height)
		if useColor {
			output = renderer.RenderImageHalfcell(resized, defaultColor)
		} else {
			gray := renderer.ToGrayscale(resized)
			output = renderer.RenderImageHalfcellGrayscale(gray)
		}
		// Remove empty lines
		lines := strings.Split(output, "\n")
		var cleanLines []string
		for _, line := range lines {
			if strings.TrimSpace(line) != "" {
				cleanLines = append(cleanLines, line)
			}
		}
		output = strings.Join(cleanLines, "\n")
	case "fullcell":
		resized = resizeImage(img, width, height)
		if useColor {
			output = renderer.RenderImageFullcell(resized, defaultColor)
		} else {
			gray := renderer.ToGrayscale(resized)
			output = renderer.RenderImageFullcellGrayscale(gray)
		}
	default:
		return "", fmt.Errorf("unsupported render type: %s (use 'halfcell' or 'fullcell')", renderType)
	}

	return output, nil
}
