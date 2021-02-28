package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(x, y int, z string) (int, string) {
	return x + y, z
}

func main() {
	var x = 2
	defer fmt.Println("It will be printed at the last")
	fmt.Println("The time is", time.Now())
	fmt.Println(x)
	fmt.Println(rand.Intn(10))
	fmt.Println(add(1, 2, "something"))
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	k := 2
	if half := k % 2; half == 1 {
		fmt.Println("old")
	} else {
		fmt.Println("young")
	}
	fmt.Println(sum2)
	a := 1
	p := &a
	fmt.Println(p)
	fmt.Println(*p)
	i := 10
	j := 20
	p = &i
	*p = *p + 1
	fmt.Println(i)
	p = &j
	*p = *p + 1
	fmt.Println(j, i)
	// fmt.Println(Split(3))
}
