package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"fmt"
)

func isBlack(c color.Color) bool {
	r, g, b, _ := c.RGBA()
	return r < 3000 && g < 3000 && b < 3000
}

func main() {
	file, err := os.Open("image.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()
	width := bounds.Dx()

	var top, bottom, left, right int
	topFound := false

	// logging cordinate
	var borderCoords []string

	// detect black square
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		blackCount := 0
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if isBlack(img.At(x, y)) {
				blackCount++
			}
		}
		if blackCount > width/2 {
			if !topFound {
				top = y
				topFound = true
			}
			bottom = y
		}
	}

	// detect left  right + log koordinat
	left = width
	right = 0
	for y := top; y <= bottom; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if isBlack(img.At(x, y)) {
				// Simpan koordinat ke log
				borderCoords = append(borderCoords, fmt.Sprintf("(%d, %d)", x, y))

				if x < left {
					left = x
				}
				if x > right {
					right = x
				}
			}
		}
	}

	// Crop image
	cropped := image.NewRGBA(image.Rect(0, 0, right-left, bottom-top))
	draw.Draw(cropped, cropped.Bounds(), img, image.Point{X: left, Y: top}, draw.Src)

	// Save output
	outFile, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, cropped)
	if err != nil {
		log.Fatal(err)
	}

	// Save border_coords
	logFile, err := os.Create("border_coords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	for _, coord := range borderCoords {
		logFile.WriteString(coord + "\n")
	}

	log.Println("Berhasil crop dan simpan log")
}