package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/misterunix/colorworks/hsl"
)

const (
	RADTODEG = 180.0 / math.Pi
	DEGTORAD = math.Pi / 180.0
)

func main() {
	// height : The height of the image.
	height := 1000
	// width : The width of the image.
	width := 1000

	// amplitude : The amplitude of the wave.
	amplitude := 0.45 * float64(width)

	for i := 1; i < 20; i++ {
		for j := 1; j < 20; j++ {
			n := float64(j)                          // n : The numerator.
			d := float64(i)                          // d : The denominator.
			drawRose(width, height, n, d, amplitude) // Draw the rhodonea curve "rose".
		}
	}

}

// drawRose : Draw the rhodonea curve "rose"
func drawRose(width, height int, n, d, amplitude float64) {

	rectboundingbox := image.Rectangle{}
	rectboundingbox.Min.X = 0
	rectboundingbox.Max.X = width
	rectboundingbox.Min.Y = 0
	rectboundingbox.Max.Y = height

	// Create a new image with the same bounds as the rectangle.
	img := image.NewRGBA(rectboundingbox)
	//c := color.White

	// bgColor : The background color of the image.
	bgColor := color.RGBA{0, 0, 0, 255}
	// Draw the background.
	draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.Point{}, draw.Src)

	// Calculate the angular frequency.
	k := n / d

	// loop over the curves by .001 increments
	for s := 0.0; s < (360.0 * d); s = s + .001 {

		hue := math.Mod(s, 360.0) // hue : The hue of the color.
		saturation := 1.0         // saturation : The saturation of the color.
		level := .5               // level : The lightness of the color.

		// Calculate the color. From the HSL color model.
		r, g, b := hsl.HSLtoRGB(hue, saturation, level)

		// cc : The color from the HSL color model in the RGB color model.
		cc := color.RGBA{r, g, b, 0xff}

		angle := s * DEGTORAD // angle : The angle of the curve.

		// Cartesian coordinates

		x := amplitude * math.Cos(angle*k) * math.Cos(angle) // x : The x-coordinate of the point.
		y := amplitude * math.Cos(angle*k) * math.Sin(angle) // y : The y-coordinate of the point.

		x1 := int(x) + (width / 2)  // x1 : The x-coordinate of the point centered on the image.
		y1 := int(y) + (height / 2) // y1 : The y-coordinate of the point centered on the image.

		img.Set(x1, y1, cc) // Set the point (x, y) to color c in the image.

	}

	//
	// Save the image to a png file.
	//

	// filename : The name of the file.
	filename := fmt.Sprintf("images/%d-%d.png", int(n), int(d))
	// f : The file descriptor.
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err) // Exit if the file cannot be created.
	}
	defer f.Close() // Close the file when the function returns.

	// Encode the image to the file.
	err = png.Encode(f, img)
	if err != nil {
		log.Fatalln(err) // Exit if the image cannot be encoded.
	}

}
