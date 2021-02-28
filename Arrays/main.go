package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 1
	a[1] = 2

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]

	fmt.Println(a, primes, s)

	//slices
	names := [4]string{
		"Daniel",
		"Max",
		"Gasly",
		"Lando",
	}
	b := names[0:2]
	c := names[1:3]
	fmt.Println(c, b)
	b[0] = "xxx"
	fmt.Println(b, c)
	fmt.Println(names)

	s = primes[:2]
	fmt.Println(s)
	s = primes[1:]
	fmt.Println(s)
	s = primes[:]
	fmt.Println(s)

	s = s[:0]
	fmt.Println(s, len(s), cap(s))
	s = s[:4]
	fmt.Println(s, len(s), cap(s))

	s = s[2:]

	fmt.Println(s, len(s), cap(s))

}

//slices are dyamic array
//slice are refrecences to arrays changing slices will effect
//the underlying array
//slice literal is an array literal without the length
//we can omit the high or low bounds

// a slice has both a length and a capacity

//length is the number of elements it contains
//cpacity of a slice is the number of elements in the underlying array

//len can be obtained by len(s) and capcity by cap(s)

// nil value of the slice has a length and capacity of 0
