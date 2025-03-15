package image_utils

import (
	"image"
	"image/draw"
	"os"
	"path/filepath"
	"strings"

	"github.com/fogleman/gg"
)

const (
	// Standard dimensions for all images
	StandardWidth  = 800
	StandardHeight = 600

	// Text styling
	FontSize     = 36
	LineSpacing  = 1.5
	MaxLineWidth = 30

	// Overlay styling
	ShadowOffset = 2
	TextPadding  = 20
)

// GenerateSingleImageWithQuote creates a motivational image with a quote overlay
func GenerateSingleImageWithQuote(imagePath, quote string) string {
	// Load and resize image to standard dimensions
	img := loadAndResizeImage(imagePath, StandardWidth, StandardHeight)
	outputPath := "output/single_quote.png"

	// Ensure the output directory exists
	createOutputDir(outputPath)

	// Create drawing context
	dc := gg.NewContextForImage(img)

	// Add semi-transparent overlay for better text readability
	dc.SetRGBA(0, 0, 0, 0.5)
	dc.DrawRectangle(0, float64(dc.Height())/2-100, float64(dc.Width()), 200)
	dc.Fill()

	// Set up text properties
	dc.SetRGB(1, 1, 1)
	LoadFont(dc, "Impact.ttf", FontSize)

	// Draw text with shadow for better readability
	wrappedQuote := wrapText(quote, MaxLineWidth)

	// Draw text shadow
	dc.SetRGB(0, 0, 0)
	dc.DrawStringWrapped(wrappedQuote, float64(dc.Width()/2)+ShadowOffset, float64(dc.Height()/2)+ShadowOffset,
		0.5, 0.5, float64(dc.Width())-TextPadding*2, LineSpacing, gg.AlignCenter)

	// Draw main text
	dc.SetRGB(1, 1, 1)
	dc.DrawStringWrapped(wrappedQuote, float64(dc.Width()/2), float64(dc.Height()/2),
		0.5, 0.5, float64(dc.Width())-TextPadding*2, LineSpacing, gg.AlignCenter)

	dc.SavePNG(outputPath)
	return outputPath
}

// GenerateWindowImageWithQuote creates a 2x2 grid of images with a quote
func GenerateWindowImageWithQuote(imagePaths []string, quote string) string {
	// Standard dimensions for the grid
	width := StandardWidth * 2
	height := StandardHeight * 2

	// Create a new canvas
	canvas := image.NewRGBA(image.Rect(0, 0, width, height))

	// Process and place each image
	for i, path := range imagePaths[:4] {
		// Determine position in grid
		x := (i % 2) * StandardWidth
		y := (i / 2) * StandardHeight

		// Load and resize image
		img := loadAndResizeImage(path, StandardWidth, StandardHeight)

		// Draw to canvas
		draw.Draw(canvas, image.Rect(x, y, x+StandardWidth, y+StandardHeight),
			img, image.Point{}, draw.Over)
	}

	// Create drawing context for adding text
	dc := gg.NewContextForImage(canvas)

	// Add semi-transparent overlay for text
	dc.SetRGBA(0, 0, 0, 0.5)
	dc.DrawRectangle(0, float64(height)/2-100, float64(width), 200)
	dc.Fill()

	// Set up text properties
	LoadFont(dc, "Impact.ttf", FontSize*1.5)

	// Wrap text to fit
	wrappedQuote := wrapText(quote, MaxLineWidth)

	// Draw text shadow
	dc.SetRGB(0, 0, 0)
	dc.DrawStringWrapped(wrappedQuote, float64(width/2)+ShadowOffset, float64(height/2)+ShadowOffset,
		0.5, 0.5, float64(width)-TextPadding*4, LineSpacing, gg.AlignCenter)

	// Draw main text
	dc.SetRGB(1, 1, 1)
	dc.DrawStringWrapped(wrappedQuote, float64(width/2), float64(height/2),
		0.5, 0.5, float64(width)-TextPadding*4, LineSpacing, gg.AlignCenter)

	outputPath := "output/window_quote.png"
	createOutputDir(outputPath)
	dc.SavePNG(outputPath)

	return outputPath
}

// loadAndResizeImage loads an image and resizes it to the specified dimensions
func loadAndResizeImage(filePath string, width, height int) image.Image {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	// Create a new context with the target dimensions
	dc := gg.NewContext(width, height)

	// Get original dimensions
	bounds := img.Bounds()
	imgWidth := bounds.Dx()
	imgHeight := bounds.Dy()

	// Calculate scaling factors
	scaleX := float64(width) / float64(imgWidth)
	scaleY := float64(height) / float64(imgHeight)

	// Use the smaller scaling factor to maintain aspect ratio
	scale := scaleX
	if scaleY < scale {
		scale = scaleY
	}

	// Calculate new dimensions
	newWidth := int(float64(imgWidth) * scale)
	newHeight := int(float64(imgHeight) * scale)

	// Draw black background
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	// Draw the image centered
	dc.DrawImageAnchored(img, width/2, height/2, 0.5, 0.5)

	return dc.Image()
}

// wrapText breaks text into appropriate line lengths
func wrapText(text string, maxChars int) string {
	words := strings.Fields(text)
	var lines []string
	currentLine := ""

	for _, word := range words {
		// Check if adding this word would exceed max chars
		if len(currentLine)+len(word)+1 > maxChars && len(currentLine) > 0 {
			lines = append(lines, currentLine)
			currentLine = word
		} else if len(currentLine) == 0 {
			currentLine = word
		} else {
			currentLine += " " + word
		}
	}

	// Add the final line
	if len(currentLine) > 0 {
		lines = append(lines, currentLine)
	}

	return strings.Join(lines, "\n")
}

// createOutputDir ensures the directory for the file exists
func createOutputDir(filePath string) {
	outputDir := filepath.Dir(filePath)
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
