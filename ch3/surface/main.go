// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"os"
)

// PlotFunc is a mathematical function of type z = f(x,y)
type PlotFunc func(float64, float64) float64

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	var function PlotFunc = f

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "sin":
			function = f
		case "saddle":
			function = saddle
		case "moguls":
			function = moguls
		default:
			fmt.Fprintf(os.Stderr, "Unsupported function: %s\n", os.Args[1])
			os.Exit(1)
		}
	}
	plot(function)
}

func plot(function PlotFunc) {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aOk := corner(i+1, j, function)
			bx, by, bOk := corner(i, j, function)
			cx, cy, cOk := corner(i, j+1, function)
			dx, dy, dOk := corner(i+1, j+1, function)

			// Don't plot invalid coordinates...
			if !(aOk && bOk && cOk && dOk) {
				continue
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, function PlotFunc) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := function(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	if !isValid(sx) || !isValid(sy) {
		return 0.0, 0.0, false
	}

	return sx, sy, true
}

func isValid(val float64) bool {
	return !math.IsInf(val, 0) && !math.IsNaN(val)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func saddle(x, y float64) float64 {
	return (x*x - y*y) / 512
}

func moguls(x, y float64) float64 {
	return (math.Sin(x) + math.Sin(y+3.0*math.Pi/4)) / 16
}

//!-
