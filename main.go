package main

import (
	"fmt"
	"math"
	"math/rand"

	gd "github.com/misterunix/cgo-gd"
	"github.com/misterunix/colorworks/hsl"
)

const (
	RADTODEG = 180.0 / math.Pi
	DEGTORAD = math.Pi / 180.0
)

type nd struct {
	N int
	D int
}

var ND []nd

func main() {

	// height : The height of the image.
	height := 1000
	// width : The width of the image.
	width := 1000

	// amplitude : The amplitude of the wave.
	amplitude := 0.45 * float64(width)

	/*
		for i := 1; i < 20; i++ {
			for j := 1; j < 20; j++ {
				n := float64(j)                          // n : The numerator.
				d := float64(i)                          // d : The denominator.
				drawRose(width, height, n, d, amplitude) // Draw the rhodonea curve "rose".
			}
		}
	*/
	ND = append(ND, nd{4, 1})
	ND = append(ND, nd{4, 2})
	drawRose(width, height, amplitude)
}

// drawRose : Draw the rhodonea curve "rose"
func drawRose(width, height int, amplitude float64) {

	l := len(ND)
	n := make([]float64, l)
	d := make([]float64, l)
	for i := 0; i < l; i++ {
		n[i] = float64(ND[i].N)
		d[i] = float64(ND[i].D)
	}

	img := gd.CreateTrueColor(width, height)

	bgColor := img.ColorAllocate(0, 0, 0)
	img.Fill(0, 0, bgColor)

	for i := 0; i < l; i++ {

		// Calculate the angular frequency.
		k := n[i] / d[i]

		// Move the color around the color wheel.
		hueOffset := rand.Float64() * 360.0

		// loop over the curves by .001 increments
		for s := 0.0; s < (360.0 * d[i]); s = s + .001 {

			hue := math.Mod(s+hueOffset, 360.0) // hue : The hue of the color.
			saturation := 1.0                   // saturation : The saturation of the color.
			level := .5                         // level : The lightness of the color.

			// Calculate the color. From the HSL color model.
			r, g, b := hsl.HSLtoRGB(hue, saturation, level)

			// cc : The color from the HSL color model in the RGB color model.
			cc := img.ColorAllocate(int(r), int(g), int(b))
			//cc := color.RGBA{r, g, b, 0xff}

			angle := s * DEGTORAD // angle : The angle of the curve.

			// Cartesian coordinates

			x := amplitude * math.Cos(angle*k) * math.Cos(angle) // x : The x-coordinate of the point.
			y := amplitude * math.Cos(angle*k) * math.Sin(angle) // y : The y-coordinate of the point.

			x1 := int(x) + (width / 2)  // x1 : The x-coordinate of the point centered on the image.
			y1 := int(y) + (height / 2) // y1 : The y-coordinate of the point centered on the image.

			img.FilledEllipse(x1, y1, 3, 3, cc)

		}
	}

	//
	// Save the image to a png file.
	//

	var filename string
	filename = "images/"
	for i := 0; i < l; i++ {
		filename += fmt.Sprintf("%d-%d.", ND[i].N, ND[i].D)
		/*
			if i < l-1 {
				filename += "-"
			}
		*/
	}

	// filename : The name of the file.
	filename += "png" // := fmt.Sprintf("images/%d-%d.png", int(n), int(d))

	img.Png(filename)

}
