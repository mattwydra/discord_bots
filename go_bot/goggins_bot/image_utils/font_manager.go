package image_utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fogleman/gg"
)

const (
	// Font directory
	FontDir = "assets/fonts"

	// Default fonts - URLs for free fonts
	ImpactFontURL    = "https://github.com/google/fonts/raw/main/apache/roboto/static/Roboto-Bold.ttf"
	SecondaryFontURL = "https://github.com/google/fonts/raw/main/ofl/oswald/static/Oswald-Bold.ttf"
)

// InitializeFonts ensures necessary fonts are available
func InitializeFonts() error {
	// Create fonts directory if it doesn't exist
	if err := os.MkdirAll(FontDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create font directory: %v", err)
	}

	// Check and download primary font
	impactPath := filepath.Join(FontDir, "Impact.ttf")
	if _, err := os.Stat(impactPath); os.IsNotExist(err) {
		fmt.Println("Downloading primary font...")
		if err := downloadFont(ImpactFontURL, impactPath); err != nil {
			return fmt.Errorf("failed to download primary font: %v", err)
		}
	}

	// Check and download secondary font
	secondaryPath := filepath.Join(FontDir, "Secondary.ttf")
	if _, err := os.Stat(secondaryPath); os.IsNotExist(err) {
		fmt.Println("Downloading secondary font...")
		if err := downloadFont(SecondaryFontURL, secondaryPath); err != nil {
			// Non-critical, just log the error
			fmt.Printf("Warning: Failed to download secondary font: %v\n", err)
		}
	}

	return nil
}

// downloadFont downloads a font from URL to the specified path
func downloadFont(url, destPath string) error {
	// Create the request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if successful
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Create the file
	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// LoadFont attempts to load a font face for the given context
// Falls back to default if the specified font can't be loaded
func LoadFont(dc *gg.Context, fontName string, size float64) error {
	fontPath := filepath.Join(FontDir, fontName)

	// Try to load the requested font
	err := dc.LoadFontFace(fontPath, size)
	if err == nil {
		return nil
	}

	// Try Impact font as backup
	impactPath := filepath.Join(FontDir, "Impact.ttf")
	err = dc.LoadFontFace(impactPath, size)
	if err == nil {
		return nil
	}

	// Set default font as last resort
	dc.SetFontFace(gg.NewFace(gg.NewDefaultFontFace(), size))
	return nil
}

// GetAvailableFonts returns a list of available font files
func GetAvailableFonts() ([]string, error) {
	var fonts []string

	files, err := os.ReadDir(FontDir)
	if err != nil {
		if os.IsNotExist(err) {
			return fonts, nil // Directory doesn't exist but not an error
		}
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && (filepath.Ext(file.Name()) == ".ttf" || filepath.Ext(file.Name()) == ".otf") {
			fonts = append(fonts, file.Name())
		}
	}

	return fonts, nil
}
