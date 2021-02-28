package main

import (
	"fmt"
	"math"
)

type abser interface {
	Abs() float64
}

type I interface {
	M()
}

type F float64

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var a abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())
	a = &v

	fmt.Println(a.Abs())

	var i I = &T{"hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	var t *T
	i = t
	describe(i)
	i.M()
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}

func describe(i I) {
	fmt.Printf("(%v, %T)", i, i)
}
