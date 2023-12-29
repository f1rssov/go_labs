package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/surface", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	widthStr := r.FormValue("width")
	heightStr := r.FormValue("height")
	//color := r.FormValue("color")

	width, _ := strconv.Atoi(widthStr)
	height, _ := strconv.Atoi(heightStr)

	if width <= 0 || height <= 0 {
		width = 600
		height = 320
	}

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, color := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64, string) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, "#000000"
	}

	color := getColor(z)

	sx := width/2 + (x-y)*xyscale
	sy := height/2 + (x+y)*xyscale - z*zscale
	return sx, sy, color
}

func f(x, y float64) float64 {
	return x*x/20 - y*y/20
}

func getColor(z float64) string {
	blue := 0
	red := 0
	if z > 0 {
		red = int(z * 100)
	} else {
		blue = int(math.Abs(z) * 100)
	}
	return fmt.Sprintf("#%02x%02x%02x", red, 0, blue)
}
