package renderer

import (
	"fmt"
	"image"
	"math"

	"image/color"
)

// converts an image to grayscale.
func ToGrayscale(img image.Image) *image.Gray {
	bounds := img.Bounds()
	gray := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			// Enhanced luminance calculation with better contrast
			lum := uint8((0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)) / 256)
			// Apply contrast enhancement
			lumFloat := float64(lum) / 255.0
			// Increase contrast using power curve
			enhanced := math.Pow(lumFloat, 0.8) * 255.0
			lum = uint8(math.Min(255, math.Max(0, enhanced)))
			gray.SetGray(x, y, color.Gray{Y: lum})
		}
	}
	return gray
}

//	renders an image using "▄" (lower half block) with ANSI truecolor codes.
//
// Each terminal row represents 2 image rows.
func RenderImageHalfcell(img image.Image, defaultColor string) string {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var output string

	// Ensure even height
	if height%2 != 0 {
		height--
	}

	for y := 0; y < height; y += 2 {
		for x := range width {
			// Upper pixel (background color)
			r1, g1, b1, a1 := img.At(x, y).RGBA()
			// Lower pixel (foreground color)
			r2, g2, b2, a2 := img.At(x, y+1).RGBA()

			// Convert from 16-bit to 8-bit
			r1b, g1b, b1b := uint8(r1>>8), uint8(g1>>8), uint8(b1>>8)
			r2b, g2b, b2b := uint8(r2>>8), uint8(g2>>8), uint8(b2>>8)

			// ANSI color codes
			var style string
			if a2 > 0 { // Lower pixel not transparent
				style += fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r2b, g2b, b2b) // Foreground
			} else if defaultColor != "" {
				style += fmt.Sprintf("\x1b[38;2;%s", defaultColor)
			}
			if a1 > 0 { // Upper pixel not transparent
				style += fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r1b, g1b, b1b) // Background
			} else if defaultColor != "" {
				style += fmt.Sprintf("\x1b[48;2;%s", defaultColor)
			}

			// Use "▄" if lower pixel is not transparent, else " "
			char := "▄"
			if a2 == 0 {
				char = " "
			}
			output += style + char + "\x1b[0m"
		}
		output += "\n"
	}
	return output
}

// renders a grayscale image using halfcell approach
func RenderImageHalfcellGrayscale(img *image.Gray) string {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var output string

	// Ensure even height
	if height%2 != 0 {
		height--
	}

	for y := 0; y < height; y += 2 {
		for x := range width {
			// upper := img.GrayAt(x, y).Y
			lower := img.GrayAt(x, y+1).Y

			// Calculate grayscale intensity for both pixels
			// upperIntensity := float64(upper) / 255.0
			lowerIntensity := float64(lower) / 255.0

			// Apply contrast enhancement to each pixel
			// upperEnhanced := math.Pow(upperIntensity, 0.7)
			lowerEnhanced := math.Pow(lowerIntensity, 0.7)

			// Use the lower pixel primarily but consider upper for context
			primaryIntensity := lowerEnhanced
			// contextIntensity := (upperEnhanced + lowerEnhanced) / 2.0

			if primaryIntensity > 0.8 {
				output += "█" // Full block
			} else if primaryIntensity > 0.65 {
				output += "▓" // Dark shade
			} else if primaryIntensity > 0.45 {
				output += "▒" // Medium shade
			} else if primaryIntensity > 0.25 {
				output += "░" // Light shade
			} else if primaryIntensity > 0.05 {
				output += "▄" // Lower half block for subtle details
			} else {
				output += " " // Empty
			}
		}
		output += "\n"
	}
	return output
}

//	renders an image using full cell blocks with ANSI truecolor codes.
//
// Each terminal cell represents 1 image pixel.
func RenderImageFullcell(img image.Image, defaultColor string) string {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var output string

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			rb, gb, bb := uint8(r>>8), uint8(g>>8), uint8(b>>8)

			var style string
			if a > 0 { // Pixel not transparent
				style = fmt.Sprintf("\x1b[48;2;%d;%d;%dm", rb, gb, bb) // Background color
			} else if defaultColor != "" {
				style = fmt.Sprintf("\x1b[48;2;%s", defaultColor)
			}

			// Use two spaces to create a square-ish cell
			output += style + "  " + "\x1b[0m"
		}
		output += "\n"
	}
	return output
}

// renders a grayscale image using fullcell approach
func RenderImageFullcellGrayscale(img *image.Gray) string {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var output string

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			gray := img.GrayAt(x, y).Y

			// Calculate grayscale intensity
			intensity := float64(gray) / 255.0
			enhanced := math.Pow(intensity, 0.7)

			// More detailed thresholds
			if enhanced > 0.8 {
				output += "██" // Full block
			} else if enhanced > 0.65 {
				output += "▓▓" // Dark shade
			} else if enhanced > 0.45 {
				output += "▒▒" // Medium shade
			} else if enhanced > 0.25 {
				output += "░░" // Light shade
			} else if enhanced > 0.05 {
				output += " " // Use single space for very subtle details
			} else {
				output += "  " // Empty
			}

		}
		output += "\n"
	}
	return output
}
