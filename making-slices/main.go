package main

import (
	"fmt"
	"strings"
)

func main() {
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)
	fmt.Println("some")

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

	var pow = []int{1, 2, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d=%d\n", i, v)
	}

}

func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d %v\n", len(x), cap(x), x)
	// fmt.Println(x, len(x), cap(x))
}

//slices can contain any type including other slices

//appending to a slice
// we can append to a slice using append function
//func append(s []T, vs ...T)[]T
