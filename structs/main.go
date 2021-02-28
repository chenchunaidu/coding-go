package main

import (
	"fmt"
)

//structs are the groups of the fields
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	v := Vertex{1, 2}
	v.X = 4
	p := &v
	p.X = 1e9
	fmt.Println(v)
	fmt.Println(v1, v2, v3, p)
}
