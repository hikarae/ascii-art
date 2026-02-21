package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg" // support for jpeg
	"image/png"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// gradient of chars from dark to light
const asciiChars = " .'`^\",:;Il!i><~+_-?][}{1)(|/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"

func main() {
	// --output--direcotry set path to output directory
	diroutput := flag.String("output-directory", ".", "path to output directory")
	fileoutput := flag.String("output-file", "output.png", "path to output file")

	// --input-directory set path to input directory default directory . --input-file set image file name default image file name input.jpg
	dirinput := flag.String("input-directory", ".", "Path to image directory")
	fileinput := flag.String("input-file", "input.jpg", "Image file name")
	flag.Parse()

	fullinputpath := filepath.Join(*dirinput, *fileinput)
	fulloutputpath := filepath.Join(*diroutput, *fileoutput)

	file, err := os.Open(fullinputpath)
	if err != nil {
		log.Fatalf("Error with open image %s: %v ", fullinputpath, err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Error with decode image: %v", err)
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

	outFile, err := os.Create(fulloutputpath)
	if err != nil {
		log.Fatalf("Error with create image: %v", err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, outImg)
	if err != nil {
		log.Fatalf("Error with save image: %v", err)
	}

	fmt.Println("Save in ", fulloutputpath)
}
