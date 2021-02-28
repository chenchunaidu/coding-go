package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var n = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m, n)

	o := make(map[string]int)
	o["Daniel riccardo"] = 7
	o["Lewis hamilton"] = 90
	o["Estaban ocon"] = 0
	fmt.Println(o)
	fmt.Println(o["Daniel riccardo"])

	v, ok := o["Lewis hamilton"]
	fmt.Println(v, ok)
	delete(o, "Estaban ocon")
	fmt.Println(o)

	sente := "some other guy"
	some := strings.Fields(sente)
	for _, word := range some {
		fmt.Println(word)
	}
}

//A map maps keys to values
//the zero value of the map is nil A nil map has no keys,nor can keys be added

//Map literals are like struct literals but the keys are required

//if toplevel name is type name we can omit it from the elements of the literal

//mutating maps
