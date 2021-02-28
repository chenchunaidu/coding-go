package main

import (
	"fmt"
)

func main() {
	i, j := 10, 20
	p := &i
	*p = *p + 1
	fmt.Println(i)
	p = &j
	*p = *p + 1
	fmt.Println(i, j)
	fmt.Println("SOMETHIN")
}
