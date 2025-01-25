package image_utils

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/fogleman/gg"
)

func GenerateSingleImageWithQuote(imagePath, quote string) string {
	img := loadImage(imagePath)
	outputPath := "output/single_quote.jpg"

	// Ensure the output directory exists
	createOutputDir(outputPath)

	dc := gg.NewContextForImage(img)
	dc.SetRGB(1, 1, 1)
	// dc.SetFontSize(32) // You can experiment with this value to make the text larger
	dc.DrawStringAnchored(quote, float64(dc.Width()/2), float64(dc.Height()/2), 0.5, 0.5)
	dc.SavePNG(outputPath)
	// dc.SetFontFace(24)

	return outputPath
}

func GenerateWindowImageWithQuote(imagePaths []string, quote string) string {
	// Load images
	img1 := loadImage(imagePaths[0])
	img2 := loadImage(imagePaths[1])
	img3 := loadImage(imagePaths[2])
	img4 := loadImage(imagePaths[3])

	// Create new canvas
	width := img1.Bounds().Dx() * 2
	height := img1.Bounds().Dy() * 2
	canvas := image.NewRGBA(image.Rect(0, 0, width, height))

	// Draw images onto canvas
	draw.Draw(canvas, image.Rect(0, 0, width/2, height/2), img1, image.Point{}, draw.Over)
	draw.Draw(canvas, image.Rect(width/2, 0, width, height/2), img2, image.Point{}, draw.Over)
	draw.Draw(canvas, image.Rect(0, height/2, width/2, height), img3, image.Point{}, draw.Over)
	draw.Draw(canvas, image.Rect(width/2, height/2, width, height), img4, image.Point{}, draw.Over)

	// Add quote to the center
	dc := gg.NewContextForImage(canvas)
	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(quote, float64(width/2), float64(height/2), 0.5, 0.5)

	outputPath := "output/window_quote.jpg"
	createOutputDir(outputPath)
	saveImage(outputPath, canvas)

	return outputPath
}

func loadImage(filePath string) image.Image {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	return img
}

func saveImage(filePath string, img image.Image) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jpeg.Encode(file, img, nil)
}

// createOutputDir ensures the directory for the file exists
func createOutputDir(filePath string) {
	outputDir := filepath.Dir(filePath)
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
