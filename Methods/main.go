package main

import (
	"fmt"
	"math"
)

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(abs(v))
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	v.Scale(10)
	fmt.Println(abs(v))
}

//we can declare a method on non-struct too
// we can declare method with pointer receivers
