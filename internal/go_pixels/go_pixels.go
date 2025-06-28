package gopixels

import (
	"fmt"
	"image"

	"go-pixels/internal/utils"

	"github.com/charmbracelet/lipgloss"
)

type PixelRenderer struct {
	Width      int
	Height     int
	Brightness int
	Style      lipgloss.Style
}

func NewPixelRenderer(width, height, brightness int, style lipgloss.Style) *PixelRenderer {
	if width <= 0 {
		width = 80 // default terminal width
	}
	if height <= 0 {
		height = 24 // default terminal height
	}
	if brightness <= 0 {
		brightness = 128 // middle brightness threshold
	}
	return &PixelRenderer{
		Width:      width,
		Height:     height,
		Brightness: brightness,
		Style:      style,
	}
}

// image to string
func (pr *PixelRenderer) FromImagePath(path string, options map[string]string) (string, error) {
	img, err := utils.LoadImage(path)
	if err != nil {
		return "", fmt.Errorf("gopixels failed to load image: %v", err)
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
		// For halfcell, resize to full width but ensure even height
		targetHeight := pr.Height
		if targetHeight%2 != 0 {
			targetHeight++
		}
		resized = utils.ResizeImage(img, pr.Width, targetHeight)
		if useColor {
			output = utils.RenderImageHalfcell(resized, defaultColor)
		} else {
			gray := utils.ToGrayscale(resized)
			output = utils.RenderImageHalfcellGrayscale(gray, pr.Brightness)
		}

	case "fullcell":
		resized = utils.ResizeImage(img, pr.Width, pr.Height)
		if useColor {
			output = utils.RenderImageFullcell(resized, defaultColor)
		} else {
			gray := utils.ToGrayscale(resized)
			output = utils.RenderImageFullcellGrayscale(gray, pr.Brightness)
		}

	default:
		return "", fmt.Errorf("unsupported render type: %s", renderType)
	}

	return output, nil
}
