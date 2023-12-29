package main

import (
	"fmt"
	"math"
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
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, color := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, string) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := saddle(x, y)
	color := getColor(z)

	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, "#000000" // Возврат нулевых координат и черного цвета для некорректных значений
	}

	sx := width/2 + (x-y)*xyscale
	sy := height/2 + (x+y)*xyscale - z*zscale
	return sx, sy, color
}

func saddle(x, y float64) float64 {
	return x*x/20 - y*y/20
}

func getColor(z float64) string {
	// Определите логику для выбора цвета на основе высоты (z) здесь
	// В данном случае, можно сделать цвет более насыщенным, если z увеличивается
	blue := 0
	red := 0
	if z > 0 {
		red = int(z * 100)
	} else {
		blue = int(math.Abs(z) * 100)
	}
	return fmt.Sprintf("#%02x%02x%02x", red, 0, blue)
}
