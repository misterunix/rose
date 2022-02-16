package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"

	"github.com/misterunix/colorworks/hsl"
)

func main() {

	width := 1000
	height := 1000

	r := image.Rectangle{}
	r.Min.X = 0
	r.Max.X = width
	r.Min.Y = 0
	r.Max.Y = height

	img := image.NewRGBA(r)
	//c := color.White

	bgColor := color.RGBA{0, 0, 0, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.Point{}, draw.Src)

	var angle float64
	var ampl float64
	var k float64

	ampl = float64(width) * .5
	//k:=2
	d := 9.0
	n := 8.0
	k = n / d

	for s := 0.0; s < (360.0 * d); s = s + .001 {

		hue := math.Mod(s, 360.0)
		saturation := 1.0
		level := .5

		r, g, b := hsl.HSLtoRGB(hue, saturation, level)

		cc := color.RGBA{r, g, b, 0xff}

		angle = s * (math.Pi / 180.0)
		x := ampl * math.Cos(angle*k) * math.Cos(angle)
		y := ampl * math.Cos(angle*k) * math.Sin(angle)

		x1 := int(x) + (width / 2)
		y1 := int(y) + (height / 2)

		img.Set(x1, y1, cc)

	}

	f, err := os.Create("test.png")
	if err != nil {
		// Handle error
	}
	defer f.Close()

	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(f, img)
	if err != nil {
		// Handle error
	}

}
