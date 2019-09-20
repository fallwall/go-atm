package main

import "fmt"

func main() {

	// to create an array a that wiill hold exactly 8 ints
	// type of elements and length are both part of array's type
	var a [5]int
	fmt.Println("emp:", a)

	// set valeu: array[index] = value
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// len(a) to return length of an array
	fmt.Println("len:", len(a))

	// to declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// array types are 1d, but you can compose tyypes to build multiple dimentional data strcutures
	var twoD [2][3]int
	for i := 0; i <2; i++ {
		for j := 0; j <3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
