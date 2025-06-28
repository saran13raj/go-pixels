package gopixels

import (
	"fmt"
	"image"
	"strings"

	"github.com/saran13raj/go-pixels/utils"
)

// FromImagePath converts an image to string representation
func FromImagePath(path string, width, height int, options map[string]string) (string, error) {
	img, err := utils.LoadImage(path)
	if err != nil {
		return "", fmt.Errorf("gopixels failed to load image: %v", err)
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

	// Parse options
	renderType := "halfcell" // default to halfcell like Python library
	useColor := true
	defaultColor := ""
	if options != nil {
		if val, ok := options["type"]; ok {
			switch val {
			case "halfcell", "fullcell", "braille", "blocks":
				renderType = val
			}
		}
		if val, ok := options["color"]; ok {
			useColor = val == "true"
		}
		if val, ok := options["default_color"]; ok {
			defaultColor = val
		}
	}

	var output string
	var resized image.Image

	switch renderType {
	case "halfcell":
		// For halfcell, use the specified height directly without adjustment
		resized = utils.ResizeImage(img, width, height)
		if useColor {
			output = utils.RenderImageHalfcell(resized, defaultColor)
		} else {
			gray := utils.ToGrayscale(resized)
			output = utils.RenderImageHalfcellGrayscale(gray)
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
		resized = utils.ResizeImage(img, width, height)
		if useColor {
			output = utils.RenderImageFullcell(resized, defaultColor)
		} else {
			gray := utils.ToGrayscale(resized)
			output = utils.RenderImageFullcellGrayscale(gray)
		}
	default:
		return "", fmt.Errorf("unsupported render type: %s", renderType)
	}

	return output, nil
}
