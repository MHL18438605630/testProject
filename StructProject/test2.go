package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q *Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type ColorPoint struct {
	point Point
	color color.RGBA
}

func main() {
	red := color.RGBA{R: 255, A: 255}
	blue := color.RGBA{B: 255, A: 255}

	p1 := ColorPoint{
		Point{
			1, 2,
		},
		red,
	}

	p2 := ColorPoint{Point{
		4, 6,
	},
		blue,
	}

	fmt.Println(p1.point.Distance(&p2.point))
}
