package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg" // support for jpeg
	"image/png"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// gradient of chars from dark to light
const asciiChars = " .:-=+*#%@"

func main() {
	file, err := os.Open("input.jpg")
	if err != nil {
		fmt.Println("Error with open image:", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error with decode image:", err)
		return
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	outImg := image.NewRGBA(bounds)
	draw.Draw(outImg, outImg.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src)

	d := &font.Drawer{
		Dst:  outImg,
		Src:  image.White,
		Face: basicfont.Face7x13,
	}

	charWidth := 7
	charHeight := 13

	for y := bounds.Min.Y; y < height; y += charHeight {
		for x := bounds.Min.X; x < width; x += charWidth {

			c := img.At(x+charWidth/2, y+charHeight/2)
			r, g, b, _ := c.RGBA()

			lum := (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 256.0

			charIdx := int(lum / 255.0 * float64(len(asciiChars)-1))

			if charIdx < 0 {
				charIdx = 0
			}
			if charIdx >= len(asciiChars) {
				charIdx = len(asciiChars) - 1
			}

			char := string(asciiChars[charIdx])

			d.Dot = fixed.P(x, y+charHeight)
			d.DrawString(char)
		}
	}

	outFile, err := os.Create("asciiart.png")
	if err != nil {
		fmt.Println("Error with create image:", err)
		return
	}
	defer outFile.Close()

	err = png.Encode(outFile, outImg)
	if err != nil {
		fmt.Println("Error with save image:", err)
		return
	}

	fmt.Println("Save in asciiart.png")
}
