package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		subpixels              = 2 // Количество подпикселей в каждом пикселе
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			var r, g, b, a uint32

			for subpy := 0; subpy < subpixels; subpy++ {
				for subpx := 0; subpx < subpixels; subpx++ {
					x := float64(px*subpixels+subpx)/float64(width*subpixels)*(xmax-xmin) + xmin
					y := float64(py*subpixels+subpy)/float64(height*subpixels)*(ymax-ymin) + ymin

					z := complex(x, y)
					col := mandelbrot(z)

					rgba := color.RGBAModel.Convert(col).(color.RGBA) // Преобразование color.Color в color.RGBA

					r += uint32(rgba.R)
					g += uint32(rgba.G)
					b += uint32(rgba.B)
					a += uint32(rgba.A)
				}
			}

			r /= uint32(subpixels * subpixels)
			g /= uint32(subpixels * subpixels)
			b /= uint32(subpixels * subpixels)
			a /= uint32(subpixels * subpixels)

			img.Set(px, py, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}

	outputFile, err := os.Create("mandelbrot_supersampled.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	png.Encode(outputFile, img)
}

func mandelbrot(z complex128) color.Color {
	const (
		iterations = 200
		contrast   = 15
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r, g, b := n*8, n*2, 255-n*8
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}
