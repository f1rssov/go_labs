package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	width, height = 800, 800
)

func main() {
	http.HandleFunc("/fractal", FractalHandler)
	http.ListenAndServe(":8080", nil)
}

func FractalHandler(w http.ResponseWriter, r *http.Request) {
	// Параметры запроса: x, y, и масштабирование
	xParam := r.URL.Query().Get("x")
	yParam := r.URL.Query().Get("y")
	scaleParam := r.URL.Query().Get("scale")

	x, _ := strconv.ParseFloat(xParam, 64)
	y, _ := strconv.ParseFloat(yParam, 64)
	scale, _ := strconv.ParseFloat(scaleParam, 64)

	img := GenerateFractal(x, y, scale)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}

func GenerateFractal(x, y, scale float64) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		yCoord := y - (float64(py)/height-0.5)*scale
		for px := 0; px < width; px++ {
			xCoord := x + (float64(px)/width-0.5)*scale
			z := complex(xCoord, yCoord)
			col := Mandelbrot(z)
			img.Set(px, py, col)
		}
	}
	return img
}

func Mandelbrot(z complex128) color.Color {
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

//http://localhost:8080/fractal?x=-0.5&y=0&scale=1.5
